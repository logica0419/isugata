package isugata

import (
	"fmt"
	"image/png"
	"net/http"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

// CorrectionLevel is the correction level of QR code
type correctionLevel string

const (
	// CorrectionLevelL is correction level L (approx 7%)
	CorrectionLevelL correctionLevel = "L"

	// CorrectionLevelM is correction level M (approx 15%)
	CorrectionLevelM correctionLevel = "M"

	// CorrectionLevelQ is correction level Q (approx 25%)
	CorrectionLevelQ correctionLevel = "Q"

	// CorrectionLevelH is correction level H (approx 30%)
	CorrectionLevelH correctionLevel = "H"
)

// WithQRCodeEqual validates if the QR code equals to the expected
func WithQRCodeEqual(size int, corrLevel correctionLevel, content string, decryptFunc func(string) (string, error)) ValidateOpt {
	return func(res *http.Response) error {
		img, err := png.Decode(res.Body)
		if err != nil {
			return fmt.Errorf("%w: %w", ErrUndecodableBody, err)
		}

		rect := img.Bounds()
		if rect.Dx() != size || rect.Dy() != size {
			return fmt.Errorf("%w: invalid QR code size: expected: %d, actual: %d", ErrInvalidBody, size, rect.Dx())
		}

		bmp, err := gozxing.NewBinaryBitmapFromImage(img)
		if err != nil {
			return fmt.Errorf("%w: %w", ErrUndecodableBody, err)
		}

		qrReader := qrcode.NewQRCodeReader()

		result, err := qrReader.Decode(bmp, nil)
		if err != nil {
			return fmt.Errorf("%w: %w", ErrUndecodableBody, err)
		}

		actualCorrectionLevel, _ := (result.GetResultMetadata()[gozxing.ResultMetadataType_ERROR_CORRECTION_LEVEL]).(string)
		if actualCorrectionLevel != string(corrLevel) {
			return fmt.Errorf("%w: invalid QR code correction level: expected: %s, actual: %s", ErrInvalidBody, corrLevel, actualCorrectionLevel)
		}

		decryptedContent, err := decryptFunc(result.String())
		if err != nil {
			return fmt.Errorf("%w: %w", ErrInvalidBody, err)
		}

		if decryptedContent != content {
			return fmt.Errorf("%w: invalid QR code content: expected: %s, actual: %s", ErrInvalidBody, content, decryptedContent)
		}

		return nil
	}
}
