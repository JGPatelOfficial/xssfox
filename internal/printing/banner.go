package printing

import "github.com/JGPatelOfficial/xssfox/pkg/model"

// Banner is XSSFox banner function
func Banner(options model.Options) {
	XSSLog("", `                                                        
               ░█▒               
             ████     ▓                    
           ▓█████  ▓██▓                  
          ████████████         ░          
        ░███████████▓          ▓░     
     ░████████████████        ▒██░    
    ▓██████████▒███████     ░█████▓░    
   ██████████████░ ████        █▓     
 ░█████▓          ░████▒       ░         XSSFox `+VERSION+`
 █████               ▓██░             
 ████                  ▓██      Powerful open-source XSS scanner       
 ███▓        ▓███████▓▒▓█░     and utility focused on automation.       
 ███▒      █████                     
 ▓███     ██████                    
 ████     ██████▒                
 ░████    ████████▒
 `, options)
}
