# HTTP Adaptive streaming tester

> This project is used to test adaptive streaming over HTTP with MPEG-DASH using DASH.js. 

Used to collect data for a report we created this project to test how DASH.js handles adaptive streaming and how it helps to provide the user with a better streaming experience. By collecting and displaying available debug data we can inspect what fragments are requested and what quality is being played. Also the buffer size is being plotted in the graph. RAW data can be extracted to be used in other softwares like `gnuplot`. There is also a webserver available to test streaming from local sources. This is good if you wish to finetune bandwidth conditions through throttling in the browser. If you wish to test without the webserver there are online sources available aswell.

![A picture of what the test site looks like](https://rtek.cloud/dash.png)
  
# Usage