# Infra

## Summary

1. [Logger](#logger)
   1. [Functionality](#functionality)
   2. [How to use](#how-to-use)

### Logger

path: `infra/config/logger.go`

##### Functionality

Tool for logging messages during program execution. Utilities:

1. Event logging
2. Debugging
3. Monitoring

##### How to use

To instantiate the logger somewhere in the application, it is necessary to use the `config` package. At the top of the file, declare the logger with file scope, considering a controller:

```go
//

package controllers

import (
 "drawfc24-api/infra/config"
 ... //imports
)

var logger = config.GetLogger("teams-controller") // Instantiating the logger

... // controller

func (pc *TeamsController) Create(w http.ResponseWriter, r *http.Request){
 ...
 logger.Error(err) // Using
 logger.Info("Look there")
}
```

> It is necessary to pass the context the logger will be executed through the function parameter.

The logger has some methods:

- Debug - Debugf
- Info - Infof
- Warn - Warnf
- Error - Errorf
