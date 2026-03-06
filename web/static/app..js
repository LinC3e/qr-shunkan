function generateQR() {

  const url = document.getElementById("url").value

  const img = document.getElementById("qr")

  img.src = `/api/qr?url=${encodeURIComponent(url)}&size=256`

}