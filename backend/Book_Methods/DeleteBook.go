package Book_Methods

import (
	"Ebook/ent/book"
	"Ebook/Lock_Methods"
	"Ebook/Structs"
	"Ebook/Token_Methods"
	"Ebook/Util_Methods"
	"context"
	"github.com/gin-gonic/gin"
	"sync"
)

// 選択本削除ハンドラ関数
func BookDeleteHandler(c *gin.Context) {
	// 変数reqをBookRegisterRequestで定義
	var req []structs.BookRegisterRequest

	//reqに取得したデータを格納、変換でエラーが起きた場合はエラーを返して終了
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(403, gin.H{"error": "変換できません"})
		return
	}
	//ゴルーチン定義
	var WaitGroup sync.WaitGroup
	//要素を1つずつ削除する
	for _, item := range req {
		WaitGroup.Add(1)
		go func(item structs.BookRegisterRequest) {
			defer WaitGroup.Done()

			//エスケープ処理
			e_Name := Util_Methods.EscapeInput(item.Name)
			e_Title := Util_Methods.EscapeInput(item.Title)
			e_Token := Util_Methods.EscapeInput(item.Token)
			//titleとuserIDからBookIDを取得
			//ユーザーNameからIDを求める
			subimiteduser, err := Util_Methods.GetUserEnt(e_Name)

			//ユーザー名が存在しない
			if err != nil {
				c.JSON(403, gin.H{"error": "invalid credentials"})
				return
			}

			//アカウントがロックされているか確認
			//あればこの場でreturnする
			existlocklist := Lock_Methods.CheckLocklist(structs.Client)

			if structs.ExistCheckfunc(subimiteduser.ID, existlocklist) {
				c.JSON(403, gin.H{"error": "invalid credentials"})
				return
			}

			//トークンを確認する
			if !Token_Methods.CheckNameToken(e_Name, e_Token) {
				c.JSON(401, gin.H{"error": "token"})
				return
			}
			//BookIDを取得する
			target_book, err := structs.Client.Book.Query().
				Where(book.UserIDEQ(subimiteduser.ID), book.TitleEQ(e_Title)).
				First(context.Background())

			if err != nil {
				//存在しない
				c.JSON(403, gin.H{"error": "invalid credentials"})
				return
			}

			//BookIDを削除
			err = structs.Client.Book.
				DeleteOneID(target_book.ID).
				Exec(context.Background())

			if err != nil {
				c.JSON(404, gin.H{"error": "削除に失敗しました。"})
				return
			}

		}(item)
	}

	//ゴルーチンの処理を待つ
	WaitGroup.Wait()
	c.JSON(200, gin.H{"book": "ok"})
}