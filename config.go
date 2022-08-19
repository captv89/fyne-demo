package main

type Config struct {
	LocalImagesFolder             string `yaml:"local_images_folder"`
	ImageObjectPath               string `yaml:"image_object_path"`
	LocalThumbsFolder             string `yaml:"local_thumbs_folder"`
	ThumbObjectPath               string `yaml:"thumb_object_path"`
	ServiceAccountKeyPath         string `yaml:"service_account_key_path"`
	ImgExifAttrName               string `yaml:"img_exif_attr_name"`
	LocalDataFilePath             string `yaml:"local_data_file_path"`
	FStorageBucketName            string `yaml:"f_storage_bucket_name"`
	ImgLinkExpiryDate             string `yaml:"img_link_expiry_date"`
	FireProjectID                 string `yaml:"fire_project_id"`
	FireDatabaseURL               string `yaml:"fire_database_url"`
	FirebaseDatabaseAccessKeyPath string `yaml:"firebase_database_access_key_path"`
	FireStoreCollectionName       string `yaml:"fire_store_collection_name"`
	FireStoreDocumentName         string `yaml:"fire_store_document_name"`
}

// Sample data for config
var config = Config{
	LocalImagesFolder:             "./images",
	ImageObjectPath:               "images",
	LocalThumbsFolder:             "./thumbs",
	ThumbObjectPath:               "thumbs",
	ServiceAccountKeyPath:         "./serviceAccountKey.json",
	ImgExifAttrName:               "exif:DateTimeOriginal",
	LocalDataFilePath:             "./data.json",
	FStorageBucketName:            "images-storage",
	ImgLinkExpiryDate:             "2020-01-01",
	FireProjectID:                 "images-project",
	FireDatabaseURL:               "https://images-project.firebaseio.com",
	FirebaseDatabaseAccessKeyPath: "./firebase_database_access_key.json",
	FireStoreCollectionName:       "images",
	FireStoreDocumentName:         "images",
}

// Update config with values from config input from user
func (c *Config) Update(config Config) {
	c.LocalImagesFolder = config.LocalImagesFolder
	c.ImageObjectPath = config.ImageObjectPath
	c.LocalThumbsFolder = config.LocalThumbsFolder
	c.ThumbObjectPath = config.ThumbObjectPath
	c.ServiceAccountKeyPath = config.ServiceAccountKeyPath
	c.ImgExifAttrName = config.ImgExifAttrName
	c.LocalDataFilePath = config.LocalDataFilePath
	c.FStorageBucketName = config.FStorageBucketName
	c.ImgLinkExpiryDate = config.ImgLinkExpiryDate
	c.FireProjectID = config.FireProjectID
	c.FireDatabaseURL = config.FireDatabaseURL
	c.FirebaseDatabaseAccessKeyPath = config.FirebaseDatabaseAccessKeyPath
	c.FireStoreCollectionName = config.FireStoreCollectionName
	c.FireStoreDocumentName = config.FireStoreDocumentName
}

