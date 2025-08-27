package main

import (
	"encoding/json" 
	"fmt"
	"html/template" 
	"log"
	"net/http" 
	"strings"
	"sync" // (Mutex)
)

// Định nghĩa một căn phòng trong trò chơi
type Room struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Exits       map[string]string `json:"exits"` // Map: hướng (north, south...) -> tên phòng đích
	Items       []string          `json:"items"` // Slice: các vật phẩm hiện có trong phòng
}

// Định nghĩa người chơi
type Player struct {
	CurrentRoom string   `json:"currentRoom"` // Phòng hiện tại của người chơi
	Inventory   []string `json:"inventory"`   // Slice: các vật phẩm người chơi đang mang
}

// Map toàn cục chứa tất cả các phòng trong trò chơi
// Key: tên phòng (string), Value: Room struct
var world = make(map[string]Room)

// Biến toàn cục đại diện cho người chơi
var player Player

// Bảo vệ trạng thái game (world và player) khỏi các race conditions
var gameMutex sync.Mutex

// Khởi tạo thế giới trò chơi và người chơi
func initGame() {
	// Khởi tạo các phòng và thêm vào world map
	world["start_room"] = Room{
		Name:        "Phòng Khởi Đầu",
		Description: "Bạn đang ở trong một căn phòng nhỏ, có vẻ như là điểm khởi đầu của một cuộc phiêu lưu. Có một cánh cửa về phía bắc. Bạn thấy một chiếc chìa khóa sáng bóng trên bàn.",
		Exits:       map[string]string{"north": "hallway"},
		Items:       []string{"chìa khóa"}, 
	}

	world["hallway"] = Room{
		Name:        "Hành Lang",
		Description: "Bạn đang ở trong một hành lang dài và tối. Có lối đi về phía nam (trở lại phòng khởi đầu) và phía đông. Một ngọn đuốc đang cháy leo lét trên tường.",
		Exits:       map[string]string{"south": "start_room", "east": "treasure_room"},
		Items:       []string{"ngọn đuốc"},
	}

	world["treasure_room"] = Room{
		Name:        "Phòng Kho Báu",
		Description: "Một căn phòng lấp lánh ánh vàng! Có rất nhiều rương kho báu ở đây. Lối ra duy nhất là về phía tây.",
		Exits:       map[string]string{"west": "hallway"},
		Items:       []string{"vàng", "ngọc bích"},
	}

	// Khởi tạo người chơi
	player = Player{
		CurrentRoom: "start_room", // Bắt đầu ở phòng "start_room"
		Inventory:   make([]string, 0), // Hành trang rỗng (slice rỗng)
	}
}

// Trả về mô tả của phòng hiện tại
func displayCurrentRoom() string {
	var output strings.Builder
	currentRoom, ok := world[player.CurrentRoom]
	if !ok {
		return "Lỗi: Phòng hiện tại không tồn tại. Vui lòng liên hệ quản trị viên."
	}

	output.WriteString(fmt.Sprintf("\n--- %s ---\n", currentRoom.Name))
	output.WriteString(currentRoom.Description + "\n")

	if len(currentRoom.Items) > 0 {
		output.WriteString("Bạn thấy: ")
		output.WriteString(strings.Join(currentRoom.Items, ", ") + ".\n")
	} else {
		output.WriteString("Không có vật phẩm nào đáng chú ý ở đây.\n")
	}

	output.WriteString("Các lối ra: ")
	exits := make([]string, 0, len(currentRoom.Exits))
	for direction := range currentRoom.Exits {
		exits = append(exits, direction)
	}
	output.WriteString(strings.Join(exits, ", ") + ".\n")

	return output.String()
}

// Xử lý lệnh của người chơi và trả về kết quả dưới dạng text
func handleCommand(command string) string {
	parts := strings.Fields(strings.ToLower(command))
	if len(parts) == 0 {
		return "" 
	}

	verb := parts[0]
	var output strings.Builder

	switch verb {
	case "go":
		if len(parts) < 2 {
			return "Đi đâu? (ví dụ: go north)"
		}
		direction := parts[1]
		currentRoom := world[player.CurrentRoom]
		if nextRoomName, ok := currentRoom.Exits[direction]; ok {
			player.CurrentRoom = nextRoomName
			output.WriteString(displayCurrentRoom()) 
		} else {
			output.WriteString(fmt.Sprintf("Bạn không thể đi về phía '%s' từ đây.\n", direction))
		}

	case "take":
		if len(parts) < 2 {
			return "Nhặt cái gì? (ví dụ: take chìa khóa)"
		}
		itemToTake := strings.Join(parts[1:], " ")
		currentRoom := world[player.CurrentRoom]
		found := false
		newRoomItems := make([]string, 0)

		for _, item := range currentRoom.Items {
			if item == itemToTake {
				player.Inventory = append(player.Inventory, item)
				output.WriteString(fmt.Sprintf("Bạn đã nhặt '%s'.\n", item))
				found = true
			} else {
				newRoomItems = append(newRoomItems, item)
			}
		}

		if found {
			currentRoom.Items = newRoomItems
			world[player.CurrentRoom] = currentRoom 
		} else {
			output.WriteString(fmt.Sprintf("Không tìm thấy '%s' ở đây.\n", itemToTake))
		}

	case "inventory", "inv":
		output.WriteString("\n--- Hành trang của bạn ---\n")
		if len(player.Inventory) == 0 {
			output.WriteString("Hành trang của bạn trống rỗng.\n")
		} else {
			for i, item := range player.Inventory {
				output.WriteString(fmt.Sprintf("%d. %s\n", i+1, item))
			}
		}

	case "look":
		output.WriteString(displayCurrentRoom())

	case "help":
		output.WriteString("\nCác lệnh có thể sử dụng:\n")
		output.WriteString("  go [hướng]  - Di chuyển theo hướng (ví dụ: go north)\n")
		output.WriteString("  take [vật phẩm] - Nhặt vật phẩm (ví dụ: take chìa khóa)\n")
		output.WriteString("  inventory / inv - Xem hành trang\n")
		output.WriteString("  look        - Mô tả lại căn phòng\n")
		output.WriteString("  quit        - Thoát trò chơi (chỉ thoát khỏi giao diện web, không tắt server, reload để chơi lại)\n")

	case "quit":
		output.WriteString("Cảm ơn bạn đã chơi! Bạn có thể đóng tab trình duyệt.\n")

	default:
		output.WriteString("Lệnh không hợp lệ. Gõ 'help' để xem các lệnh.\n")
	}
	return output.String()
}

// --- Các hàm xử lý HTTP ---
type CommandRequest struct {
	Command string `json:"command"`
}

type CommandResponse struct {
	Output string `json:"output"`
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// Xử lý các lệnh game từ client
func handleGameCommand(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Chỉ chấp nhận phương thức POST", http.StatusMethodNotAllowed)
		return
	}

	var req CommandRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Lỗi đọc yêu cầu JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Bảo vệ trạng thái game bằng Mutex
	gameMutex.Lock()
	defer gameMutex.Unlock()

	// Xử lý lệnh và lấy output
	gameOutput := handleCommand(req.Command)

	// Chuẩn bị và gửi phản hồi JSON
	resp := CommandResponse{Output: gameOutput}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	initGame() 

	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/command", handleGameCommand)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server đang chạy tại http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}