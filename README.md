README.md
=========

RSSGator  
A fast, minimal CLI utility for aggregating and reading RSS feeds.

Installation
------------

### From source

```bash
git clone https://github.com/iceblade92/RSSGator.git
cd RSSGator
go build -o rssgator ./cmd/rssgator
sudo mv rssgator /usr/local/bin
```


Quick start
-----------

1. **Add feeds**  
   ```bash
   rssgator add https://example.com/feed.xml
   ```

2. **Update and read**  
   ```bash
   rssgator update        # fetch latest posts
   rssgator list          # show all feeds
   rssgator read 5        # read 5 most recent posts
   ```

3. **Mark as read**  
   ```bash
   rssgator mark 42       # mark post #42 read
   ```

Configuration
-------------

Config and data live in `~/.config/rssgator`.  
Edit `config.toml` if you need to tweak fetch intervals or output style.

Defenetly used AI for this read me did not bother to make one...