package pdfcpu

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"github.com/pkg/errors"
)

var (
	pad = []byte{
		0x28, 0xBF, 0x4E, 0x5E, 0x4E, 0x75, 0x8A, 0x41, 0x64, 0x00, 0x4E, 0x56, 0xFF, 0xFA, 0x01, 0x08,
		0x2E, 0x2E, 0x00, 0xB6, 0xD0, 0x68, 0x3E, 0x80, 0x2F, 0x0C, 0xA9, 0xFE, 0x64, 0x53, 0x69, 0x7A,
	}

	nullPad32 = make([]byte, 32)

	perm = map[model.CommandMode]struct{ extract, modify int }{
		model.VALIDATE:                {0, 0},
		model.LISTINFO:                {0, 0},
		model.OPTIMIZE:                {0, 0},
		model.SPLIT:                   {1, 0},
		model.SPLITBYPAGENR:           {1, 0},
		model.MERGECREATE:             {0, 0},
		model.MERGECREATEZIP:          {0, 0},
		model.MERGEAPPEND:             {0, 0},
		model.EXTRACTIMAGES:           {1, 0},
		model.EXTRACTFONTS:            {1, 0},
		model.EXTRACTPAGES:            {1, 0},
		model.EXTRACTCONTENT:          {1, 0},
		model.EXTRACTMETADATA:         {1, 0},
		model.TRIM:                    {0, 1},
		model.LISTATTACHMENTS:         {0, 0},
		model.EXTRACTATTACHMENTS:      {1, 0},
		model.ADDATTACHMENTS:          {0, 1},
		model.ADDATTACHMENTSPORTFOLIO: {0, 1},
		model.REMOVEATTACHMENTS:       {0, 1},
		model.LISTPERMISSIONS:         {0, 0},
		model.SETPERMISSIONS:          {0, 0},
		model.ADDWATERMARKS:           {0, 1},
		model.REMOVEWATERMARKS:        {0, 1},
		model.IMPORTIMAGES:            {0, 1},
		model.INSERTPAGESBEFORE:       {0, 1},
		model.INSERTPAGESAFTER:        {0, 1},
		model.REMOVEPAGES:             {0, 1},
		model.LISTKEYWORDS:            {0, 0},
		model.ADDKEYWORDS:             {0, 1},
		model.REMOVEKEYWORDS:          {0, 1},
		model.LISTPROPERTIES:          {0, 0},
		model.ADDPROPERTIES:           {0, 1},
		model.REMOVEPROPERTIES:        {0, 1},
		model.COLLECT:                 {1, 0},
		model.CROP:                    {0, 1},
		model.LISTBOXES:               {0, 0},
		model.ADDBOXES:                {0, 1},
		model.REMOVEBOXES:             {0, 1},
		model.LISTANNOTATIONS:         {0, 1},
		model.ADDANNOTATIONS:          {0, 1},
		model.REMOVEANNOTATIONS:       {0, 1},
		model.ROTATE:                  {0, 1},
		model.NUP:                     {0, 1},
		model.BOOKLET:                 {0, 1},
		model.LISTBOOKMARKS:           {0, 0},
		model.ADDBOOKMARKS:            {0, 1},
		model.REMOVEBOOKMARKS:         {0, 1},
		model.IMPORTBOOKMARKS:         {0, 1},
		model.EXPORTBOOKMARKS:         {0, 1},
		model.LISTIMAGES:              {0, 1},
		model.UPDATEIMAGES:            {0, 1},
		model.CREATE:                  {0, 0},
		model.DUMP:                    {0, 1},
		model.LISTFORMFIELDS:          {0, 0},
		model.REMOVEFORMFIELDS:        {0, 1},
		model.LOCKFORMFIELDS:          {0, 1},
		model.UNLOCKFORMFIELDS:        {0, 1},
		model.RESETFORMFIELDS:         {0, 1},
		model.EXPORTFORMFIELDS:        {0, 1},
		model.FILLFORMFIELDS:          {0, 1},
		model.LISTPAGELAYOUT:          {0, 1},
		model.SETPAGELAYOUT:           {0, 1},
		model.RESETPAGELAYOUT:         {0, 1},
		model.LISTPAGEMODE:            {0, 1},
		model.SETPAGEMODE:             {0, 1},
		model.RESETPAGEMODE:           {0, 1},
		model.LISTVIEWERPREFERENCES:   {0, 1},
		model.SETVIEWERPREFERENCES:    {0, 1},
		model.RESETVIEWERPREFERENCES:  {0, 1},
		model.ZOOM:                    {0, 1},
	}

	ErrUnknownEncryption = errors.New("pdfcpu: unknown encryption")
)

func newEncryptDict(pdf20, needAES bool, keyLength int, permissions int16) types.Dict {
	_ = "STUB: not implemented"
	return *new(types.Dict)
}

func encKey(userpw string, e *model.Enc) (key []byte) { _ = "STUB: not implemented"; return nil }

func validateUserPassword(ctx *model.Context) (ok bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func key(ownerpw, userpw string, r, l int) (key []byte) { _ = "STUB: not implemented"; return nil }

func o(ctx *model.Context) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func u(ctx *model.Context) (u []byte, key []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func validationSalt(bb []byte) []byte { _ = "STUB: not implemented"; return nil }

func keySalt(bb []byte) []byte { _ = "STUB: not implemented"; return nil }

func decryptOE(ctx *model.Context, opw []byte) error { _ = "STUB: not implemented"; return nil }

func validateOwnerPasswordAES256(ctx *model.Context) (ok bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func decryptUE(ctx *model.Context, upw []byte) error { _ = "STUB: not implemented"; return nil }

func validateUserPasswordAES256(ctx *model.Context) (ok bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func processInput(input string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func hashRev6(input, pw, U []byte) ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func validateOwnerPasswordAES256Rev6(ctx *model.Context) (ok bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func validateUserPasswordAES256Rev6(ctx *model.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func validateOwnerPassword(ctx *model.Context) (ok bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func perms(p int) (list []string) { _ = "STUB: not implemented"; return nil }

func PermissionsList(p int) (list []string) { _ = "STUB: not implemented"; return nil }

func Permissions(ctx *model.Context) (list []string) { _ = "STUB: not implemented"; return nil }

func validatePermissions(ctx *model.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func writePermissions(ctx *model.Context, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func logP(enc *model.Enc) { _ = "STUB: not implemented"; return }

func maskExtract(mode model.CommandMode, secHandlerRev int) int {
	_ = "STUB: not implemented"
	return 0
}

func maskModify(mode model.CommandMode, secHandlerRev int) int { _ = "STUB: not implemented"; return 0 }

func hasNeededPermissions(mode model.CommandMode, enc *model.Enc) bool {
	_ = "STUB: not implemented"
	return false
}

func getR(ctx *model.Context, d types.Dict) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func validateAlgorithm(ctx *model.Context) (ok bool) { _ = "STUB: not implemented"; return false }

func validateAES256Parameters(d types.Dict, r int) (oe, ue, perms []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func validateOAndU(ctx *model.Context, d types.Dict, r int) (o, u []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func checkCFLengthV2(len int, pdf20, relaxed bool) error { _ = "STUB: not implemented"; return nil }

func checkCFLengthAESV2(len int, pdf20, relaxed bool) error { _ = "STUB: not implemented"; return nil }

func validateCFLength(len *int, cfm *string, pdf20, relaxed bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCryptFilterRecipients(ctx *model.Context, d types.Dict, cfm *string) error {
	_ = "STUB: not implemented"
	return nil
}

func checkCryptFilterCFM(cfm string, v int) error { _ = "STUB: not implemented"; return nil }

func validateCryptFilter(ctx *model.Context, d types.Dict, v int, pubKeySecHandler, relaxed bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func locateCFEntry(ctx *model.Context, d types.Dict, v int, key string, pubKeySecHandler, relaxed bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func validateStmf(ctx *model.Context, d, cfDict types.Dict, v int, pubKeySecHandler, relaxed bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStrf(ctx *model.Context, d, cfDict types.Dict, v int, pubKeySecHandler, relaxed bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateEFF(ctx *model.Context, d, cfDict types.Dict, v int, pubKeySecHandler, relaxed bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCryptFilters(ctx *model.Context, d types.Dict, v int, pubKeySecHandler bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateEncryptFilter(d types.Dict) (string, error) { _ = "STUB: not implemented"; return "", nil }

func validateEncryptSubFilter(d types.Dict, pubKeySecHandler bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func validateEncryptV(d types.Dict) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func validateEncryptLength(d types.Dict, v int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func validatePubKeySecHandler(ctx *model.Context, d types.Dict, pubKeySecHandler bool, subFilter string) error {
	_ = "STUB: not implemented"
	return nil
}

func supportedEncryption(ctx *model.Context, d types.Dict) (*model.Enc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decryptKey(objNumber, generation int, key []byte, aes bool) []byte {
	_ = "STUB: not implemented"
	return nil
}

func encryptBytes(b []byte, objNr, genNr int, encKey []byte, needAES bool, r int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decryptBytes(b []byte, objNr, genNr int, encKey []byte, needAES bool, r int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func applyRC4CipherBytes(b []byte, objNr, genNr int, key []byte, needAES bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encrypt(m map[string]types.Object, k string, v types.Object, objNr, genNr int, key []byte, needAES bool, r int) error {
	_ = "STUB: not implemented"
	return nil
}

func encryptDict(d types.Dict, objNr, genNr int, key []byte, needAES bool, r int) error {
	_ = "STUB: not implemented"
	return nil
}

func encryptStringLiteral(sl types.StringLiteral, objNr, genNr int, key []byte, needAES bool, r int) (*types.StringLiteral, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decryptStringLiteral(sl types.StringLiteral, objNr, genNr int, key []byte, needAES bool, r int) (*types.StringLiteral, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encryptHexLiteral(hl types.HexLiteral, objNr, genNr int, key []byte, needAES bool, r int) (*types.HexLiteral, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decryptHexLiteral(hl types.HexLiteral, objNr, genNr int, key []byte, needAES bool, r int) (*types.HexLiteral, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encryptDeepObject(objIn types.Object, objNr, genNr int, key []byte, needAES bool, r int) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func decryptDict(d types.Dict, objNr, genNr int, key []byte, needAES bool, r int) error {
	_ = "STUB: not implemented"
	return nil
}

func decryptDeepObject(objIn types.Object, objNr, genNr int, key []byte, needAES bool, r int) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func encryptStream(buf []byte, objNr, genNr int, encKey []byte, needAES bool, r int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decryptStream(buf []byte, objNr, genNr int, encKey []byte, needAES bool, r int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func applyRC4Bytes(buf, key []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func encryptAESBytes(b, key []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func decryptAESBytes(b, key []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func fileID(ctx *model.Context) (types.HexLiteral, error) {
	_ = "STUB: not implemented"
	return *new(types.HexLiteral), nil
}

func calcFileEncKey(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }

func calcOAndUAES256(ctx *model.Context, d types.Dict) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func calcOAndUAES256Rev6(ctx *model.Context, d types.Dict) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func calcOAndU(ctx *model.Context, d types.Dict) (err error) { _ = "STUB: not implemented"; return nil }
