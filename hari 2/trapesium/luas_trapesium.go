package main
import "fmt"
func main() {
  var sa, sb, t,luas, phi float32
  phi = 0.5
  fmt.Print("Masukkan panjang sisi sejajar atas , bawah , tinggi: ")
  fmt.Scanln(&sa)
  fmt.Scanln(&sb)
  fmt.Scanln(&t)

  luas = phi * (sa + sb) * t
  fmt.Println("Luas trapesium adalah ", luas)
}