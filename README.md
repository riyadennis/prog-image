# prog-image
API to upload,download and convert  images
To run the api for testing:
 make run

Use curl commands below:

To Upload:

curl -X POST \
  http://localhost:8080/upload/ \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
	"uri": "https://vignette.wikia.nocookie.net/creepypasta/images/e/e2/Anime-Girl-With-Silver-Hair-And-Purple-Eyes-HD-Wallpaper.jpg/revision/latest?cb=20140120061808",
	"type": "png"
}'


To Convert:

curl -X GET \
  'http://localhost:8080/images/57c714d7-5e0a-11e8-86dc-985aeb8c5470?type=gif'


To Download :

curl -X GET \
  http://localhost:8080/57c714d7-5e0a-11e8-86dc-985aeb8c5470.gif