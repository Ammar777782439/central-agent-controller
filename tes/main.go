package main

import (
	"fmt"
	"os"
)

func main() {
	// محاولة فتح ملف
	file, err := os.ReadFile("agent/main.go")

	// تحقق إذا في خطأ
	if err != nil {
		// في خطأ حصل
		fmt.Println("❌ فشل فتح الملف:", err)
		return // نوقف البرنامج
	}

	// لو ما في خطأ
	fmt.Println("✅ تم فتح الملف بنجاح:", file)
}
