## Photo of the day application

Traverses all subdirectories at a given location and randomly selects a photo (file) and emails it.

Used to send an email of the day by pointing it at my Photos for MacOS original photos directory and setting up a cron job to run once per day.

```Usage of ./random-photo:
  -d string
        Path to photos directory
  -i int
        SMTP Port, defaults to 587 (default 587)
  -m string
        outbound smtp server, defaults to smtp.gmail.com (default "smtp.gmail.com")
  -p string
        SMTP password
  -u string
        SMTP username
  -w string
        Addresses to send to```