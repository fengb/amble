# amble

Quickly dumps a bunch of requests for comparison.

1. Create `headers`
   1. Pop open networks tab
   2. Right click the request -> Copy -> Copy Request Headers
   3. `$ pbpaste > headers`

2. Create `endpoints`
   ```bash
   $ echo "/" >> endpoints
   $ echo "/foo/bar" >> endpoints
   ```

3. Generate into directory
   ```bash
   $ amble headers endpoints before
   # git checkout after_refactor
   $ amble headers endpoints after
   ```

4. Compare
   ```bash
   $ diff before after
   ```
