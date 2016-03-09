Application allow to configure redirect after some service migration.

```usage
Usage of ./goredirect:

  -dport int
    	redirect to port (default:80) (default 80)
  -sport int
    	redirect from port (default:8080) (default 8080)
  -template string
    	template to redirect (default "http://[%HOST%]:[%PORT%]/[%PATH%]")
  -rcode int
      	responce code (default:301) (default 301)
```usage
      	
