package telegram

// PassportData contains information about Telegram Passport data shared with the bot by the user.
//
// https://core.telegram.org/bots/api#passportdata
type PassportData struct {
	Data        []*EncryptedPassportElement `json:"data"`
	Credentials *EncryptedCredentials       `json:"credentials"`
}

// PassportFile represents a file uploaded to Telegram Passport. Currently all Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
//
// https://core.telegram.org/bots/api#passportfile
type PassportFile struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FileDate     int    `json:"file_date"`
}

// EncryptedPassportElement contains information about documents or other Telegram Passport elements shared with the bot by the user.
//
// https://core.telegram.org/bots/api#encryptedpassportelement
type EncryptedPassportElement struct {
	Type        string          `json:"type"`
	Data        string          `json:"data,omitempty"`         // Optional.
	PhoneNumber string          `json:"phone_number,omitempty"` // Optional.
	Email       string          `json:"email,omitempty"`        // Optional.
	Files       []*PassportFile `json:"files,omitempty"`        // Optional.
	FrontSide   *PassportFile   `json:"front_side,omitempty"`   // Optional.
	ReverseSide *PassportFile   `json:"reverse_side,omitempty"` // Optional.
	Selfie      *PassportFile   `json:"selfie,omitempty"`       // Optional.
	Translation []*PassportFile `json:"translation,omitempty"`  // Optional.
	Hash        string          `json:"hash"`
}

// EncryptedCredentials contains data required for decrypting and authenticating EncryptedPassportElement.
// See the Telegram Passport Documentation for a complete description of the data decryption and authentication processes.
//
// https://core.telegram.org/bots/api#encryptedcredentials
type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

// PassportElementError represents an error in the Telegram Passport element which was submitted that should be resolved by the user. It should be one of:
//
// PassportElementErrorDataField
// PassportElementErrorFrontSide
// PassportElementErrorReverseSide
// PassportElementErrorSelfie
// PassportElementErrorFile
// PassportElementErrorFiles
// PassportElementErrorTranslationFile
// PassportElementErrorTranslationFiles
// PassportElementErrorUnspecified
//
// https://core.telegram.org/bots/api#passportelementerror
// TODO: should PassportElementError be a type?

// PassportElementErrorDataField represents an issue in one of the data fields that was provided by the user.
// The error is considered resolved when the field's value changes.
//
// https://core.telegram.org/bots/api#passportelementerrordatafield
type PassportElementErrorDataField struct {
	Source    string `json:"source"`
	Type      string `json:"type"`
	FieldName string `json:"field_name"`
	DataHash  string `json:"data_hash"`
	Message   string `json:"message"`
}

// PassportElementErrorFrontSide represents an issue with the front side of a document.
// The error is considered resolved when the file with the front side of the document changes.
//
// https://core.telegram.org/bots/api#passportelementerrorfrontside
type PassportElementErrorFrontSide struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

// PassportElementErrorReverseSide represents an issue with the reverse side of a document.
// The error is considered resolved when the file with reverse side of the document changes.
//
// https://core.telegram.org/bots/api#passportelementerrorreverseside
type PassportElementErrorReverseSide struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

// PassportElementErrorSelfie represents an issue with the selfie with a document.
// The error is considered resolved when the file with the selfie changes.
//
// https://core.telegram.org/bots/api#passportelementerrorselfie
type PassportElementErrorSelfie struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

// PassportElementErrorFile represents an issue with a document scan.
// The error is considered resolved when the file with the document scan changes.
//
// https://core.telegram.org/bots/api#passportelementerrorfile
type PassportElementErrorFile struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

// PassportElementErrorFiles represents an issue with a list of scans.
// The error is considered resolved when the list of files containing the scans changes.
//
// https://core.telegram.org/bots/api#passportelementerrorfiles
type PassportElementErrorFiles struct {
	Source     string   `json:"source"`
	Type       string   `json:"type"`
	FileHashes []string `json:"file_hashes"`
	Message    string   `json:"message"`
}

// PassportElementErrorTranslationFile represents an issue with one of the files that constitute the translation of a document.
// The error is considered resolved when the file changes.
//
// https://core.telegram.org/bots/api#passportelementerrortranslationfile
type PassportElementErrorTranslationFile struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

// PassportElementErrorTranslationFiles represents an issue with the translated version of a document.
// The error is considered resolved when a file with the document translation change.
//
// https://core.telegram.org/bots/api#passportelementerrortranslationfiles
type PassportElementErrorTranslationFiles struct {
	Source     string   `json:"source"`
	Type       string   `json:"type"`
	FileHashes []string `json:"file_hashes"`
	Message    string   `json:"message"`
}

// PassportElementErrorUnspecified represents an issue in an unspecified place.
// The error is considered resolved when new data is added.
//
// https://core.telegram.org/bots/api#passportelementerrorunspecified
type PassportElementErrorUnspecified struct {
	Source      string `json:"source"`
	Type        string `json:"type"`
	ElementHash string `json:"element_hash"`
	Message     string `json:"message"`
}
