package oracle

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"time"
)

// ABitOfEverything represents a row from 'a_bit_of_everything.a_bit_of_everything'.
type ABitOfEverything struct {
	ABool                               bool            `json:"a_bool"`                                    // a_bool
	ABoolNullable                       sql.NullBool    `json:"a_bool_nullable"`                           // a_bool_nullable
	ABlob                               []byte          `json:"a_blob"`                                    // a_blob
	ABlobNullable                       []byte          `json:"a_blob_nullable"`                           // a_blob_nullable
	AChar                               string          `json:"a_char"`                                    // a_char
	ACharNullable                       sql.NullString  `json:"a_char_nullable"`                           // a_char_nullable
	ACharacter                          string          `json:"a_character"`                               // a_character
	ACharacterNullable                  sql.NullString  `json:"a_character_nullable"`                      // a_character_nullable
	AClob                               string          `json:"a_clob"`                                    // a_clob
	AClobNullable                       sql.NullString  `json:"a_clob_nullable"`                           // a_clob_nullable
	ADate                               time.Time       `json:"a_date"`                                    // a_date
	ADateNullable                       sql.NullTime    `json:"a_date_nullable"`                           // a_date_nullable
	ADoublePrecision                    float64         `json:"a_double_precision"`                        // a_double_precision
	ADoublePrecisionNullable            sql.NullFloat64 `json:"a_double_precision_nullable"`               // a_double_precision_nullable
	ADecimal                            int             `json:"a_decimal"`                                 // a_decimal
	ADecimalNullable                    sql.NullInt64   `json:"a_decimal_nullable"`                        // a_decimal_nullable
	AFloat                              float64         `json:"a_float"`                                   // a_float
	AFloatNullable                      sql.NullFloat64 `json:"a_float_nullable"`                          // a_float_nullable
	AInt                                int             `json:"a_int"`                                     // a_int
	AIntNullable                        sql.NullInt64   `json:"a_int_nullable"`                            // a_int_nullable
	AInteger                            int             `json:"a_integer"`                                 // a_integer
	AIntegerNullable                    sql.NullInt64   `json:"a_integer_nullable"`                        // a_integer_nullable
	ALongRaw                            []byte          `json:"a_long_raw"`                                // a_long_raw
	ANchar                              string          `json:"a_nchar"`                                   // a_nchar
	ANcharNullable                      sql.NullString  `json:"a_nchar_nullable"`                          // a_nchar_nullable
	ANclob                              string          `json:"a_nclob"`                                   // a_nclob
	ANclobNullable                      sql.NullString  `json:"a_nclob_nullable"`                          // a_nclob_nullable
	ANumber                             int64           `json:"a_number"`                                  // a_number
	ANumberNullable                     sql.NullInt64   `json:"a_number_nullable"`                         // a_number_nullable
	ANumeric                            int             `json:"a_numeric"`                                 // a_numeric
	ANumericNullable                    sql.NullInt64   `json:"a_numeric_nullable"`                        // a_numeric_nullable
	ANvarchar2                          string          `json:"a_nvarchar2"`                               // a_nvarchar2
	ANvarchar2Nullable                  sql.NullString  `json:"a_nvarchar2_nullable"`                      // a_nvarchar2_nullable
	ARaw                                []byte          `json:"a_raw"`                                     // a_raw
	ARawNullable                        []byte          `json:"a_raw_nullable"`                            // a_raw_nullable
	AReal                               float64         `json:"a_real"`                                    // a_real
	ARealNullable                       sql.NullFloat64 `json:"a_real_nullable"`                           // a_real_nullable
	ARowid                              string          `json:"a_rowid"`                                   // a_rowid
	ARowidNullable                      sql.NullString  `json:"a_rowid_nullable"`                          // a_rowid_nullable
	ASmallint                           int             `json:"a_smallint"`                                // a_smallint
	ASmallintNullable                   sql.NullInt64   `json:"a_smallint_nullable"`                       // a_smallint_nullable
	ATimestamp                          time.Time       `json:"a_timestamp"`                               // a_timestamp
	ATimestampNullable                  sql.NullTime    `json:"a_timestamp_nullable"`                      // a_timestamp_nullable
	ATimestampWithLocalTimeZone         time.Time       `json:"a_timestamp_with_local_time_zone"`          // a_timestamp_with_local_time_zone
	ATimestampWithLocalTimeZoneNullable sql.NullTime    `json:"a_timestamp_with_local_time_zone_nullable"` // a_timestamp_with_local_time_zone_nullable
	ATimestampWithTimeZone              time.Time       `json:"a_timestamp_with_time_zone"`                // a_timestamp_with_time_zone
	ATimestampWithTimeZoneNullable      sql.NullTime    `json:"a_timestamp_with_time_zone_nullable"`       // a_timestamp_with_time_zone_nullable
	AVarchar                            string          `json:"a_varchar"`                                 // a_varchar
	AVarcharNullable                    sql.NullString  `json:"a_varchar_nullable"`                        // a_varchar_nullable
	AVarchar2                           string          `json:"a_varchar2"`                                // a_varchar2
	AVarchar2Nullable                   sql.NullString  `json:"a_varchar2_nullable"`                       // a_varchar2_nullable
	AXmltype                            []byte          `json:"a_xmltype"`                                 // a_xmltype
	AXmltypeNullable                    []byte          `json:"a_xmltype_nullable"`                        // a_xmltype_nullable
}
