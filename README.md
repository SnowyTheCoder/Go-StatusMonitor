![Go-StatusMonitor Logo][logo]<br>
A simple status monitor that checks the status of servers.<br>

![GitHub](https://img.shields.io/github/license/Go-StatusMonitor/Go-StatusMonitor?style=for-the-badge)
<br><br>
# What it does
Go-StatusMonitor checks a list of servers and outputs the status in to a MySQL table.<br>
If any servers are down it will send a email alerting you that servers are down!<br>
It out puts into a MySQL table so you could display the status of servers in tools like Grafana.<br>
<br><br>
# How do I setup the MySQL table?
It's simple just make a new Table a bit like this:<br>
```sql
CREATE TABLE GoStatusMonitor (
Server_Name VARCHAR(200) NOT NULL,
Status VARCHAR(30) NOT NULL,
);
```
<br><br>
## FAQ:
#### How do I compile and use this?
Run "go build go-statusmonitor.go" after changing everything you want to change,<br>
Then run the compiled version!<br>
<br>
#### Can I use this to check the status of minecraft servers?
Yes.
<br>
#### Can I use this to check the status of my website?
Yes.
<br>
#### Can I download and use this for private / commercial use?
Yes, you can do anything you want with code / program as long as you do not remove the copyright message and/or the license.
<br>
#### Can I use the logo?
Yes, you can use it if you do not claim it as your own.
<br><br>
# Contributors
#### [Joshua Sing @joshuasing](https://github.com/joshuasing)
#### [Joel Sing @4a6f656c](https://github.com/4a6f656c)

<details>
  <summary>License</summary>
BSD 2-Clause License

Copyright (c) 2019, Joshua Sing
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

</details>

[logo]: https://raw.githubusercontent.com/Go-StatusMonitor/Go-StatusMonitor/master/logo/Go-StatusMonitor%20Logo%20Cropped.png "Go-StatusMonitor Logo"
