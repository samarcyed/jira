package jiradata

/////////////////////////////////////////////////////////////////////////
// This Code is Generated by SlipScheme Project:
// https://github.com/coryb/slipscheme
//
// Generated with command:
// slipscheme -dir jiradata -pkg jiradata -overwrite schemas/CurrentUser.json
/////////////////////////////////////////////////////////////////////////
//                            DO NOT EDIT                              //
/////////////////////////////////////////////////////////////////////////

// LoginInfo defined from schema:
// {
//   "title": "Login Info",
//   "type": "object",
//   "properties": {
//     "failedLoginCount": {
//       "title": "failedLoginCount",
//       "type": "integer"
//     },
//     "lastFailedLoginTime": {
//       "title": "lastFailedLoginTime",
//       "type": "string"
//     },
//     "loginCount": {
//       "title": "loginCount",
//       "type": "integer"
//     },
//     "previousLoginTime": {
//       "title": "previousLoginTime",
//       "type": "string"
//     }
//   }
// }
type LoginInfo struct {
	FailedLoginCount    int    `json:"failedLoginCount,omitempty" yaml:"failedLoginCount,omitempty"`
	LastFailedLoginTime string `json:"lastFailedLoginTime,omitempty" yaml:"lastFailedLoginTime,omitempty"`
	LoginCount          int    `json:"loginCount,omitempty" yaml:"loginCount,omitempty"`
	PreviousLoginTime   string `json:"previousLoginTime,omitempty" yaml:"previousLoginTime,omitempty"`
}
