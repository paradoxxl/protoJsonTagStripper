###### README

This project was created to strip the json 'omitempty' tags from protobuf compiled code.
Only primitives such as int64, string, ... and enums are considered as structs are embedded via pointer and should be omitted if they are null.

**Usage**

Either specify an folder and specify if it should be traversed recursively: 
`protoJsonTagStripper -folder <FolderPath> -recursive <bool>`

Or specify a single file which should be processed:
`protoJsonTagStripper -file <FilePath>`

Only files with `.pb.go` file endings will be processes.