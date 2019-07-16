Classic URL Shortener Service


Create a URL shortener library (e.g., bitly). Given a long URL, your program should return a shortened URL. Ensure the following:

a. for a long URL X, your program should always return the shortened URL x

b. the short URLs your program returns, must not follow any pattern; successive calls to your program should return very different short URLs 

c. for two different long URLs X and Y, your program should (ideally) always return two different shortened URLs x and y

d. your program must be able to accept long URLs provided both through console as well as from a file of newline separated long URLs

e. Provide the stats, number of clicks on a particular url, which are popular urls


...........................................

project1 is the main routes file. 

imported created packages (bulk,custom_new,packages,token_url,post)

Bulk package has function related to  bulk uploads from bulk_upload.html page, storing data in table links and saving file uploaded in saved

Custom_new package has custom short link generation 

Post has function related to short link generation 

Packages has token_create file which has function to create random token which is mapped to large link.

token_url has url_token function which redirect short url to large url 
............................................

