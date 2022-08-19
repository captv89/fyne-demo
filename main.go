package main

import (
	"bytes"
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	adminGUI()
}

var doImageAction, doDataAction, doUploadImages, doUploadData bool

// Text Grid to outpput the results
var (
	textGrid = widget.NewTextGrid()
	b        bytes.Buffer
)

var adminContainer = container.NewVScroll(textGrid)

func updateTerminalOutput(msg string) {
	b.WriteString(msg + "\n")
	textGrid.SetText(b.String())
	adminContainer.ScrollToBottom()
}

func adminGUI() { // Intialize a new application
	myApp := app.New()
	myWindow := myApp.NewWindow("Demo")
	myWindow.Resize(fyne.NewSize(800, 500))
	adminContainer.SetMinSize(fyne.NewSize(800, 200))

	seperator := widget.NewSeparator()

	//  Tab 1: Admin
	// Define contents in the canvas

	sub1dataActioncheck := widget.NewCheck("Perform Sub1 Data Action and save it to Json", func(value bool) {
		updateTerminalOutput(fmt.Sprintf("Perform Sub1 Data Action and save it to Json: %t", value))
		log.Println("Sub1 Data Action Check set to", value)
	})
	sub1dataActioncheck.Hide()


	sub2dataActioncheck := widget.NewCheck("Perform Sub2 Data Action and save it to Json", func(value bool) {
		updateTerminalOutput(fmt.Sprintf("Perform Sub2 Data Action and save it to Json: %t", value))
		log.Println("Sub2 Data Action Check set to", value)
	})
	sub2dataActioncheck.Hide()

	sub1imageActioncheck := widget.NewCheck("Perform Sub1 Image Action and save blurhash to exif", func(value bool) {
		updateTerminalOutput(fmt.Sprintf("Perform Sub1 Image Action and save blurhash to exif: %t", value))
		log.Println("Sub1 Image Action Check set to", value)
	})
	sub1imageActioncheck.Hide()

	sub2imageActioncheck := widget.NewCheck("Perform Sub2 Image Action and save blurhash to exif", func(value bool) {
		updateTerminalOutput(fmt.Sprintf("Perform Sub2 Image Action and save blurhash to exif: %t", value))
		log.Println("Sub2 Image Action Check set to", value)
	})
	sub2imageActioncheck.Hide()

	imageActioncheck := widget.NewCheck("Perform Image Action and save blurhash to exif", func(value bool) {
		updateTerminalOutput(fmt.Sprintf("Perform Image Action and save blurhash to exif: %t", value))
		log.Println("Perform Image Action and save blurhash to exif: ", value)
		if value {
			sub1imageActioncheck.Show()
			sub2imageActioncheck.Show()
		} else {
			sub1imageActioncheck.Hide()
			sub2imageActioncheck.Hide()
		}
	})

	uploadImagescheck := widget.NewCheck("Automatically upload Images to firestore", func(value bool) {
		updateTerminalOutput(fmt.Sprintf("Automatically upload Images to firestore: %t", value))
		log.Println("Upload Images Check set to", value)
	})

	dataActioncheck := widget.NewCheck("Perform Data Action and save it to Json", func(value bool) {
		updateTerminalOutput(fmt.Sprintf("Perform Data Action and save it to Json: %t", value))
		log.Println("Data Action Check set to", value)
		if value {
			sub1dataActioncheck.Show()
			sub2dataActioncheck.Show()
		} else {
			sub1dataActioncheck.Hide()
			sub2dataActioncheck.Hide()
		}
	})

	uploadDataCheck := widget.NewCheck("Automatically upload Data to firestore", func(value bool) {
		updateTerminalOutput(fmt.Sprintf("Automatically upload Data to firestore: %t", value))
		log.Println("Upload Data Check set to", value)
	})

	myIndendedBox := container.New(
		layout.NewVBoxLayout(),
		container.NewCenter(sub1dataActioncheck),
		layout.NewSpacer(),
		container.NewCenter(sub2dataActioncheck),
	)

	

	// Button to start the admin controller
	startButton := widget.NewButton("Start", func() {
		log.Println("Start Button Pressed")
		doImageAction = imageActioncheck.Checked
		doDataAction = dataActioncheck.Checked
		doUploadImages = uploadImagescheck.Checked
		doUploadData = uploadDataCheck.Checked
		log.Println("Image Action Checked:", doImageAction)
		log.Println("Data Action Checked:", doDataAction)
		log.Println("Upload Images Checked:", doUploadImages)
		log.Println("Upload Data Checked:", doUploadData)
		app.New().SendNotification(fyne.NewNotification("Magic Started", "Start Button Pressed"))
		adminController()
	})

	//  Tab 2: Config input form

	// Declare config values with binding
	local_images_folder_value := binding.NewString()
	image_object_path_value := binding.NewString()
	local_thumbs_folder_value := binding.NewString()
	thumb_object_path_value := binding.NewString()
	service_account_key_path_value := binding.NewString()
	img_exif_attr_name_value := binding.NewString()
	local_data_file_path_value := binding.NewString()
	f_storage_bucket_name_value := binding.NewString()
	img_link_expiry_date_value := binding.NewString()
	fire_project_id_value := binding.NewString()
	fire_database_url_value := binding.NewString()
	firebase_database_access_key_path_value := binding.NewString()
	fire_store_collection_name_value := binding.NewString()
	fire_store_document_name_value := binding.NewString()

	// Update config values
	_ = local_images_folder_value.Set(config.LocalImagesFolder)
	_ = image_object_path_value.Set(config.ImageObjectPath)
	_ = local_thumbs_folder_value.Set(config.LocalThumbsFolder)
	_ = thumb_object_path_value.Set(config.ThumbObjectPath)
	_ = service_account_key_path_value.Set(config.ServiceAccountKeyPath)
	_ = img_exif_attr_name_value.Set(config.ImgExifAttrName)
	_ = local_data_file_path_value.Set(config.LocalDataFilePath)
	_ = f_storage_bucket_name_value.Set(config.FStorageBucketName)
	_ = img_link_expiry_date_value.Set(config.ImgLinkExpiryDate)
	_ = fire_project_id_value.Set(config.FireProjectID)
	_ = fire_database_url_value.Set(config.FireDatabaseURL)
	_ = firebase_database_access_key_path_value.Set(config.FirebaseDatabaseAccessKeyPath)
	_ = fire_store_collection_name_value.Set(config.FireStoreCollectionName)
	_ = fire_store_document_name_value.Set(config.FireStoreDocumentName)

	// Create widget for each input field
	local_images_folder := widget.NewEntryWithData(local_images_folder_value)
	image_object_path := widget.NewEntryWithData(image_object_path_value)
	local_thumbs_folder := widget.NewEntryWithData(local_thumbs_folder_value)
	thumb_object_path := widget.NewEntryWithData(thumb_object_path_value)
	service_account_key_path := widget.NewEntryWithData(service_account_key_path_value)
	img_exif_attr_name := widget.NewEntryWithData(img_exif_attr_name_value)
	local_data_file_path := widget.NewEntryWithData(local_data_file_path_value)
	f_storage_bucket_name := widget.NewEntryWithData(f_storage_bucket_name_value)
	img_link_expiry_date := widget.NewEntryWithData(img_link_expiry_date_value)
	fire_project_id := widget.NewEntryWithData(fire_project_id_value)
	fire_database_url := widget.NewEntryWithData(fire_database_url_value)
	firebase_database_access_key_path := widget.NewEntryWithData(firebase_database_access_key_path_value)
	fire_store_collection_name := widget.NewEntryWithData(fire_store_collection_name_value)
	fire_store_document_name := widget.NewEntryWithData(fire_store_document_name_value)

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Local Images Folder Path", Widget: local_images_folder},
			{Text: "Firebase Storage Image Object Path", Widget: image_object_path},
			{Text: "Local Thumbs Folder Path", Widget: local_thumbs_folder},
			{Text: "Firebase Storage Thumbs Object Path", Widget: thumb_object_path},
			{Text: "Firebase Storage Access Key(json) Path", Widget: service_account_key_path},
			{Text: "Image Exif Attribute Name", Widget: img_exif_attr_name},
			{Text: "Local Data Output File(json) Path", Widget: local_data_file_path},
			{Text: "Firebase Storage Bucket Name", Widget: f_storage_bucket_name},
			{Text: "Image Link Expiry Date", Widget: img_link_expiry_date},
			{Text: "Firebase Project ID", Widget: fire_project_id},
			{Text: "Firebase Database URL", Widget: fire_database_url},
			{Text: "Firebase Database Access Key(json) Path", Widget: firebase_database_access_key_path},
			{Text: "Firebase Store Collection Name", Widget: fire_store_collection_name},
			{Text: "Firebase Store Document Name", Widget: fire_store_document_name},
		},
		OnSubmit: func() { // optional, handle form submission
			// Update config values
			config.LocalImagesFolder, _ = local_images_folder_value.Get()
			config.ImageObjectPath, _ = image_object_path_value.Get()
			config.LocalThumbsFolder, _ = local_thumbs_folder_value.Get()
			config.ThumbObjectPath, _ = thumb_object_path_value.Get()
			config.ServiceAccountKeyPath, _ = service_account_key_path_value.Get()
			config.ImgExifAttrName, _ = img_exif_attr_name_value.Get()
			config.LocalDataFilePath, _ = local_data_file_path_value.Get()
			config.FStorageBucketName, _ = f_storage_bucket_name_value.Get()
			config.ImgLinkExpiryDate, _ = img_link_expiry_date_value.Get()
			config.FireProjectID, _ = fire_project_id_value.Get()
			config.FireDatabaseURL, _ = fire_database_url_value.Get()
			config.FirebaseDatabaseAccessKeyPath, _ = firebase_database_access_key_path_value.Get()
			config.FireStoreCollectionName, _ = fire_store_collection_name_value.Get()
			config.FireStoreDocumentName, _ = fire_store_document_name_value.Get()
			log.Println("Config updated", config)
		},
		SubmitText: "Save",
	}

	//  Config Tab Setup
	configContent := container.NewVScroll(form)
	configTabItem := container.NewTabItemWithIcon("Config", theme.InfoIcon(), configContent)

	//  Admin Tab Setup
	adminContent := container.NewVBox(dataActioncheck, myIndendedBox,imageActioncheck, sub1imageActioncheck , sub2imageActioncheck, uploadImagescheck, uploadDataCheck, startButton, seperator, adminContainer)
	adminTabItem := container.NewTabItemWithIcon("Admin", theme.HomeIcon(), adminContent)
	tabs := container.NewAppTabs(adminTabItem, configTabItem)

	//  Create two app tabs

	// Set contents + Show and Run
	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}

func adminController() {
	log.Println("Admin Controller")
	if doImageAction {
		log.Println("Performing Image Action")
		imageAction()
	}
	if doUploadImages {
		log.Println("Uploading Images")
		uploadImages()
	}
	if doDataAction {
		log.Println("Performing Data Action")
		dataAction()
	}
	if doUploadData {
		log.Println("Uploading Data")
		uploadData()
	}
}

func dataAction() {
	log.Println("Data Action")
}

func imageAction() {
	log.Println("Image Action")
}

func uploadImages() {
	log.Println("Upload Images")
}

func uploadData() {
	log.Println("Upload Data")
}

func sub1dataAction() {
	log.Println("Sub1 Data Action")
}

func sub2dataAction() {
	log.Println("Sub2 Data Action")
}

func sub1imageAction() {
	log.Println("Sub1 Image Action")
}

func sub2imageAction() {
	log.Println("Sub2 Image Action")
}
