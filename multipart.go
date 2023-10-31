package booksapi

import (
	"books/gen/books"
	"io"
	"io/ioutil"
	"mime/multipart"
)

// BooksUploadImageDecoderFunc implements the multipart decoder for service
// "books" endpoint "uploadImage". The decoder must populate the argument p
// after encoding.
func BooksUploadImageDecoderFunc(mr *multipart.Reader, p **books.UploadImagePayload) error {
	// Create a new UploadImagePayload
	payload := new(books.UploadImagePayload)

	// Iterate over all parts
	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// If part is "image", read all and assign to payload.Image
		if part.FormName() == "image" {
			imgData, err := ioutil.ReadAll(part)
			if err != nil {
				return err
			}
			payload.Image = imgData
		}

		// If part is "content_type", assign to payload.ContentType
		if part.FormName() == "content_type" {
			payload.ContentType = part.Header.Get("Content-Type")
		}
	}

	// Assign payload to p
	*p = payload

	return nil
}

// BooksUploadImageEncoderFunc implements the multipart encoder for service
// "books" endpoint "uploadImage".
func BooksUploadImageEncoderFunc(mw *multipart.Writer, p *books.UploadImagePayload) error {
	// Create a new form file for "image"
	fw, err := mw.CreateFormFile("image", "image")
	if err != nil {
		return err
	}

	// Write the image data to the form file
	_, err = fw.Write(p.Image)
	if err != nil {
		return err
	}

	// Add "content_type" to the form
	err = mw.WriteField("content_type", p.ContentType)
	if err != nil {
		return err
	}

	return nil
}
