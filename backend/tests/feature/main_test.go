package feature

import (
	"corporation-site/tests/common/setup"
	"os"
	"testing"
)

// main_test.go の修正
func TestMain(m *testing.M) {
	err := setup.SetupTestServer()
	if err != nil {
		panic(err)
	}

	code := m.Run()

	// テストスイート終了時にクリーンアップを実行
	// if err := setup.TearDownTestDatabase(); err != nil {
	// 	fmt.Printf("Failed to tear down test database: %v\n", err)
	// }
	// setup.TearDownTestServer()

	os.Exit(code)
}

// setupTest は各テストの前に呼び出されるヘルパー関数
func setupTest(t *testing.T, tables ...string) {
	// TokenManagerをクリア
	// middleware.GetGlobalTokenManager().ClearAll()

	// 特定のテーブルが指定された場合はそれらのみをクリア
	// 指定がない場合は全テーブルをクリア
	var err error
	if len(tables) > 0 {
		err = setup.ResetTestDatabase(tables...)
	} else {
		err = setup.TearDownTestDatabase()
	}

	if err != nil {
		t.Fatalf("Failed to reset test database: %v", err)
	}
}
