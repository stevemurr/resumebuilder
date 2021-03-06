# Resumebuilder

This is an experiment in using toml to define a resume, then render that toml config with latex.  This decouples the view from the model and makes versioning resumes simpler (hopefully).

# Installation

You'll need docker.

# Usage

Run the server by either building or running `go run main.go`  

`http http://localhost:9001/api/toml/pdf < config.toml >> out.pdf`  

# Config

```toml
[settings]
scaling = "0.92"
cvstyle = "classic"
cvcolor = "blue"

[personal]
firstName = "Steven"
lastName = "Murr"
phoneNumber = "408.5555555"
github = "stevemurr"
email = "stevemurr@g.com"
address = "555 Lincoln Ct."
cityAndState = "San Jose, CA"
title = "Software Engineer / IoT Prototyping"

[skills]
proficient = ["Go", "Python", "NodeJS", "Bash"]
familiar = ["Swift", "Java", "C"]
frameworks = ["React", "Angular"]
libraries = ["Circle CI", "Git", "GNU Make", "jQuery"]
devices = ["Raspberry PI", "Banana PI", "Arduino", "ESP8266"]

[experience]
    [experience.microsoft]
        from = "Feb 2016"
        to = "Present"
        title = "Software Engineer / Cortana"
        company = "Microsoft"
        subtitle = ["Python", "Go", "React", "Docker", "MongoDB", "BoltDB", "Postgres"]
        points = [
            "Worked closely with third parties to develop and deploy custom speech platform software, build custom TTS voice fonts and explore new use cases for Speaker Recognition, Audio Segmentation and Spectral Repair. ",
            "Developed web portal that aggregates audio, Excel Files and font building data to allow users to better understand training data used to develop TTS fonts.",
            "Developed a web tool for studio vendors to ensure spectral consistency and allow feedback on specific features such as F0, RMS and speaking rate.",
            "Prototyped a technique to combine multiple domain specific voice fonts to improve Cortana VQ in a specific domain.",
            "Prototyped a plugin framework to allow quick modification of audio plugins on ALSA sinks for use in embedded devices.",
            "Extended existing VX work to improve user understanding of long lists. ",
            "Deployed tools on Window, Linux and Mac using Docker."
        ]
    [experience.namco]
        from = "Apr 2008"
        to = "Jul 2010"
        title = "QA Engineer"
        company = "Namco Networks"
        subtitle = ["Java"]
        points = [
            "Built and tested mobile games, fixing minor issues such as resolution device targeting, localization, and setting define flags.",
            "Received top department awards for achieving highest pass rate and consistently reporting high severity defects.",
            "Memorized multiple carrier and internal standards, authored test plans, and became adept with Sea Pines? Test Tracker Pro.",
            "Worked in close liaison with programmers and producers, acting as the bridge between engineering, QA, and production."
        ]

[education]
    [education.cs]
        to = "Ongoing"
        from = ""
        institution = "West Valley"
        degree = "B.S. in Computer Science"
        note = "60\\% Complete"
    [education.ongoing]
        to = "Ongoing"
        from = ""
        institution = "EdX Coursera"
        degree = "Self-Directed Learning"
        note = "Attended InterSpeech 2016, Harvards CS50 C programming course, Googles Python course, Berkeleys CS 162 Operating Systems course, and Stanfords Programming Paradigms course. "

[extracarricular]
    [extracarricular.gliderlabs]
        from = "Oct 2016"
        to = "Present"
        title = "Open Source Innovation Lab"
        company = "Glider Labs"
        subtitle = ["Go"]
        points = [
            "Collective of hackers looking to make the world more programmable.",
            "Contributed to the development of Go programming educational materials.",
            "Prototyped IoT sensor-based devices.",
            "\\url{http://gliderlabs.com}"
        ]

    [extracarricular.alsapluginframework]
        from = "Jun 2017"
        to = "Present"
        title = "Alsa Plugin Framework"
        company = "Personal Project"
        subtitle = ["ALSA", "Python", "Bash"]
        points = [
            "Developed to simplify usage of LADSPA audio plugins in ALSA.",
            "\\url{https://github.com/stevemurr/alsa-plugin-framework}"
        ]

    [extracarricular.oxylus]
        from = "Apr 2017"
        to = "Present"
        title = "Agriculture Automation"
        company = "Oxylus"
        subtitle = ["Go", "BoltDB", "Docker", "Kubernetes"]
        points = [
            "Developed an event system that works with off-the-shelf IoT devices.",
            "\\url{https://github.com/stevemurr/oxylus-events}"
        ]

    [extracarricular.securerouter]
        from = "Jun 2017"
        to = "Present"
        title = "Secure Router"
        company = "Personal Project"
        subtitle = ["BPI-R1", "Go", "OpenVPN"]
        points = [
            "Devloping router targeted at consumers which provides DNS level ad blocking and personal VPN. ",
            "Currently working on automating VPN setup.  "
        ]

    [extracarricular.radiomonitoring]
        from = "Jul 2014"
        to = "Present"
        title = "Internet Radio Monitoring Service"
        company = "Personal Project"
        subtitle = ["Python", "MySQL"]
        points = [
            "Built an online radio monitoring platform using audio fingerprinting to track song plays.",
            "Deployed on Google App Engine."
        ]

    [extracarricular.lecturerecord]
        from = "Jul 2015"
        to = "Aug 2015"
        title = "Raspberry PI Based Lecture Recording Device"
        company = "Personal Project"
        subtitle = ["Python"]
        points = [
            "Created a lecture-recording device using a Raspberry Pi that automates the video capture, audio recording, and file upload to the cloud. Designed as a cost effective solution for students to easily access and archive lectures. ",
            "A demo of it in action is available on YouTube at \\url{https://youtu.be/6QNrnzQUXQ0}"
        ]
```