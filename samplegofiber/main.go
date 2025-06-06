package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// สร้างแอปพลิเคชัน Fiber
	//create instance for fiber
	//basic route
	app := fiber.New()
	// กำหนดเส้นทางสำหรับการเข้าถึงหน้าแรก
	app.Get("/", func(c *fiber.Ctx) error {
		// ส่งข้อความ "Hello, World!" กลับไปยังผู้ใช้
		return c.SendString("Hello, World!")
	})
	app.Get("/about", func(c *fiber.Ctx) error {
		// ส่งข้อความ "Hello, World!" กลับไปยังผู้ใช้
		return c.SendString("About Page")
	})

	//create route
	//method get
	//pointer can get request/response c *fiber.Ctx
	app.Get("/contact", func(c *fiber.Ctx) error { // /contact is paramiter
		// ส่งข้อความ "Hello, World!" กลับไปยังผู้ใช้
		return c.SendString("Contact Page 999")
	})

	//route แบบมี raramiter
	app.Get("/user/:id", func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ "id" จาก URL
		id := c.Params("id")
		// ส่งค่าพารามิเตอร์ "id" กลับไปยังผู้ใช้
		return c.SendString("User ID: " + id)
	})

	// การกำหนด route แบบพารามิเตอร์หลายตัว
	// --------------------------------
	// กำหนดเส้นทางที่มีพารามิเตอร์หลายตัว
	// http://localhost:3000/user/123/profile
	//optional route put ? in line ex: /:profile?
	app.Get("/user/:id/:profile?", func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ "id" และ "profile" จาก URL
		id := c.Params("id")
		profile := c.Params("profile")
		// ส่งค่าพารามิเตอร์ "id" และ "profile" กลับไปยังผู้ใช้
		return c.SendString("User ID: " + id + ", Profile: " + profile)
	})

	// การกำหนด route แบบรับ query string
	// --------------------------------
	// กำหนดเส้นทางที่รับ query string
	// ตัวอย่าง URL: http://localhost:3000/search?q=golang
	app.Get("/search", func(c *fiber.Ctx) error {
		// ดึงค่าจาก query string "q"
		query := c.Query("q")
		// ส่งค่าจาก query string "q" กลับไปยังผู้ใช้
		return c.SendString("Search Query: " + query)
	})

	// การกำหนด route แบบรับ query string หลายค่า
	// --------------------------------
	// กำหนดเส้นทางที่รับ query string หลายค่า
	// ตัวอย่าง URL: http://localhost:3000/search/multiple?q=golang&page=2
	app.Get("/search/multiple", func(c *fiber.Ctx) error {
		// ดึงค่าจาก query string "q" และ "page"
		query := c.Query("q")
		page := c.Query("page")
		// ส่งค่าจาก query string "q" และ "page" กลับไปยังผู้ใช้
		return c.SendString("Search Query: " + query + ", Page: " + page)
	})

	// การกำหนด route แบบ wildcard
	// --------------------------------
	// กำหนดเส้นทางที่ใช้ wildcard เพื่อจับทุกเส้นทางที่ไม่ตรงกับเส้นทางอื่น
	// ตัวอย่าง URL: http://localhost:3000/wildcard/some/path/here
	app.Get("/wildcard/*", func(c *fiber.Ctx) error {
		// ดึงค่าจาก wildcard path
		wildcardPath := c.Params("*")
		// ส่งค่าจาก wildcard path กลับไปยังผู้ใช้
		return c.SendString("Wildcard Path: " + wildcardPath)
	})

	// การกำหนด route บังคับให้เป็นตัวเลข
	// --------------------------------
	// กำหนดเส้นทางที่มีพารามิเตอร์ที่ต้องเป็นตัวเลข
	// ตัวอย่าง URL: http://localhost:3000/customer/1234
	// app.Get("/customer/:id<int>", func(c *fiber.Ctx) error {
	//app.Get("/customer/:id<min(5)>", func(c *fiber.Ctx) error {
	app.Get("/customer/:id<min(100);maxLen(5)>", func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ "id" ที่ต้องเป็นตัวเลข
		id := c.Params("id")
		// ส่งค่าพารามิเตอร์ "id" กลับไปยังผู้ใช้
		return c.SendString("Customer ID: " + id)
	})

	// การกำหนด route แบบกำหนด พารามิเตอร์หลายตัวด้วย regex
	// --------------------------------
	// กำหนดเส้นทางที่มีพารามิเตอร์หลายตัว โดยใช้ regex เพื่อกำหนดรูปแบบ
	// http://localhost:3000/product/12345
	// รับเฉพาะตัวเลข 5 หลัก
	app.Get(`/product/:id<regex(\d{5})>`, func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ "id" ที่ต้องเป็นตัวเลข 5 หลัก
		id := c.Params("id")
		// ส่งค่าพารามิเตอร์ "id" กลับไปยังผู้ใช้
		return c.SendString("Product ID: " + id)
	})

	// การกำหนด route แบบรับพารามิเตอร์ที่เป็นตัวอักษรและตัวเลข
	// --------------------------------
	// 2024-08-27 (Must match regular expression)
	// ตัวอย่าง URL: http://localhost:3000/item/2024-08-27
	app.Get(`/item/:id<regex(\d{4}-\d{2}-\d{2})>`, func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ "id" จาก URL
		id := c.Params("id")
		// ส่งค่าพารามิเตอร์กลับไปยังผู้ใช้
		return c.SendString("Item ID: " + id)
	})

	// ---------------------------------------------
	// การกำหนด route แบบรับ HTTP Method ต่างๆ
	// รับค่า GET, POST, PUT, DELETE จาก Request

	// กำหนด struct Person สำหรับการรับข้อมูล JSON
	type Person struct {
		ID    int    `json:"id"`    // กำหนด ID เป็นตัวเลข
		Name  string `json:"name"`  // กำหนด Name เป็นตัวอักษร
		Email string `json:"email"` // กำหนด Email เป็นตัวอักษร
	}
	// ตัวแปรสำหรับเก็บข้อมูลผู้ใช้ โดยใช้ slice ของ struct Person
	var people []Person = []Person{
		{
			ID:    1,
			Name:  "John Doe",
			Email: "john@email.com",
		},
	}

	// GET: ดึงข้อมูลผู้ใช้ทั้งหมด
	app.Get("/person", func(c *fiber.Ctx) error {
		// ส่งข้อมูลผู้ใช้ทั้งหมดกลับไปยังผู้ใช้ในรูปแบบ JSON
		return c.JSON(fiber.Map{
			"message": "All person retrieved successfully",
			"data":    people,
			"count":   len(people),
		})
	})

	// POST: สร้างผู้ใช้ใหม่
	// ตัวอย่าง URL: http://localhost:3000/person
	// ต้องส่งข้อมูล JSON ใน body ของ request (payload)
	// {
	//   "id": 2,
	//   "name": "Jane Doe",
	//   "email": "jane@email.com"
	// }
	app.Post("/person", func(c *fiber.Ctx) error {
		// สร้างตัวแปรสำหรับเก็บข้อมูลผู้ใช้ใหม่
		person := new(Person)
		// แปลงข้อมูล JSON ที่ส่งมาใน body ของ request เป็น struct Person
		if err := c.BodyParser(person); err != nil {
			// ถ้าเกิดข้อผิดพลาดในการแปลงข้อมูล JSON ให้ส่งข้อความผิดพลาดกลับไป
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request data",
				"error":   err.Error(),
			})
		}

		// เพิ่มผู้ใช้ใหม่เข้าไปใน slice ของผู้ใช้
		people = append(people, *person)

		// ส่งข้อความยืนยันการสร้างผู้ใช้ใหม่กลับไปยังผู้ใช้
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Person created successfully",
			"data":    person,
			"count":   len(people),
		})
	})

	// GET: ดึงข้อมูลผู้ใช้ตาม ID
	app.Get("/person/:id", func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ "id" จาก URL
		id := c.Params("id")

		// แปลงค่าพารามิเตอร์ "id" เป็นตัวเลข
		idInt, err := strconv.Atoi(id)
		// ถ้าแปลงไม่ได้ ให้ส่งข้อความผิดพลาดกลับไป
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid ID format",
			})
		}

		// ค้นหาผู้ใช้ที่มี ID ตรงกับค่าที่ส่งมา
		for _, person := range people {
			if person.ID == idInt {
				// ถ้าพบผู้ใช้ ให้ส่งข้อมูลผู้ใช้กลับไป
				return c.JSON(fiber.Map{
					"message": "Person retrieved successfully",
					"data":    person,
				})
			}
		}

		// ถ้าไม่พบผู้ใช้ ให้ส่งข้อความผิดพลาดกลับไป
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Person not found",
			"data":    nil,
		})
	})

	// PUT: อัปเดตข้อมูลผู้ใช้ตาม ID
	// ตัวอย่าง URL: http://localhost:3000/person/1
	// ต้องส่งข้อมูล JSON ใน body ของ request (payload)
	// {
	//   "name": "John Smith",
	//   "email": "smith@email.com",
	// }
	app.Put("/person/:id", func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ "id" จาก URL
		id := c.Params("id")
		// แปลงค่าพารามิเตอร์ "id" เป็นตัวเลข
		idInt, err := strconv.Atoi(id)
		// ถ้าแปลงไม่ได้ ให้ส่งข้อความผิดพลาดกลับไป
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid ID format",
			})
		}
		// สร้างตัวแปรสำหรับเก็บข้อมูลผู้ใช้ใหม่
		person := new(Person)
		// แปลงข้อมูล JSON ที่ส่งมาใน body ของ request เป็น struct Person
		if err := c.BodyParser(person); err != nil {
			// ถ้าเกิดข้อผิดพลาดในการแปลงข้อมูล JSON ให้ส่งข้อความผิดพลาดกลับไป
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request data",
				"error":   err.Error(),
			})
		}
		// ค้นหาผู้ใช้ที่มี ID ตรงกับค่าที่ส่งมา
		for i, p := range people {
			if p.ID == idInt {
				// ถ้าพบผู้ใช้ ให้ปรับปรุงข้อมูลของผู้ใช้
				people[i].Name = person.Name
				people[i].Email = person.Email
				// ส่งข้อความยืนยันการอัปเดตข้อมูลผู้ใช้กลับไปยังผู้ใช้
				return c.JSON(fiber.Map{
					"message": "Person updated successfully",
					"data":    people[i],
				})
			}
		}
		// ถ้าไม่พบผู้ใช้ ให้ส่งข้อความผิดพลาดกลับไป
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Person not found",
			"data":    nil,
		})
	})

	// DELETE: ลบข้อมูลผู้ใช้ตาม ID
	// ตัวอย่าง URL: http://localhost:3000/person/1
	app.Delete("/person/:id", func(c *fiber.Ctx) error {
		// ดึงค่าพารามิเตอร์ "id" จาก URL
		id := c.Params("id")
		// แปลงค่าพารามิเตอร์ "id" เป็นตัวเลข
		idInt, err := strconv.Atoi(id)
		// ถ้าแปลงไม่ได้ ให้ส่งข้อความผิดพลาดกลับไป
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid ID format",
			})
		}
		// ค้นหาผู้ใช้ที่มี ID ตรงกับค่าที่ส่งมา
		for i, person := range people {
			if person.ID == idInt {
				// ถ้าพบผู้ใช้ ให้ลบข้อมูลผู้ใช้ใน slice
				// การลบข้อมูลใน slice:
				// 1. ใช้ append() เพื่อสร้าง slice ใหม่ที่ไม่มีข้อมูลที่ต้องการลบ
				// 2. people[:i] = ข้อมูลก่อนตำแหน่ง i
				// 3. people[i+1:] = ข้อมูลหลังตำแหน่ง i
				// 4. append() จะรวมข้อมูลทั้งสองส่วนเข้าด้วยกัน
				// 5. people = slice ใหม่ที่ไม่มีข้อมูลที่ต้องการลบ
				people = append(people[:i], people[i+1:]...)
				// ส่งข้อความยืนยันการลบผู้ใช้กลับไปยังผู้ใช้
				return c.JSON(fiber.Map{
					"message": "Person deleted successfully",
					"data":    person,
					"count":   len(people),
				})
			}
		}
		// ถ้าไม่พบผู้ใช้ ให้ส่งข้อความผิดพลาดกลับไป
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Person not found",
			"data":    nil,
		})
	})

	//ตัวแปรคือสิ่งที่เก็บใน memory
	// รันแอปพลิเคชันบนพอร์ต 3000
	app.Listen(":3000")
}

//hot reload is air
