// Before using any of the functions declared in config package of your application,
// [LoadConfig] function must be called to load the configuration details from
// application.yml file. Otherwise this might throw an Null Pointer Exception.
//
// All the files prefixed with internal_ are for internal usage to this package
// and must not export any function, variable, struct or interface. All files exporting
// certain functions must be prefixed with external_ for consistency.
package config
