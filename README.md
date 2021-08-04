<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->



<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">MySQL Server README</h3>

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <aAbout The Project</a>
      <ul>
        <li>About The Project</li>
      </ul>
    </li>
    <li>
      <a>Getting Started</a>
      <ul>
        <li>Prerequisites</a></li>
        <li>Docker-ENVs</a></li>
        <li>Installation</a></li>
      </ul>
    </li>
    <li>Usage</a></li>
    <li>Contact</a></li>
    <li>Acknowledgements</a></li>
  </ol>
</details>


<!-- ABOUT THE PROJECT -->
# About The Project

We created this project to 

[![Product Name Screen Shot][product-screenshot]](https://example.com)


<!-- GETTING STARTED -->
# Getting Started

In order to obtain a local copy of this project, please follow these simple steps:

1. Clone the following repo: https://github.com/ahujashivani/MySQL-Server.git
2. Install Docker (Desktop Application)
3. Complete the following terminal commands listed in the "installation" section

# Prerequisites

Please make sure to have Docker downloaded as a Desktop application and have the recently cloned directory opened in terminal. Port 8080 and 8000 should remain open until the server starts running.

<!-- Docker Image -->
# Docker-ENVs

The following Environment variables have already been added to the Docker-Compose file. These are placed here for your reference. If you are altering the database name, please make sure to alter the DB_DATABASE environment variable.

  ```sh
  DB_USER: "root"
  ```
  ```sh
  DB_PASSWORD: "rootpassword"
  ```
  ```sh
  DB_HOST: "mysql_db_container:3306"
  ```
  ```sh
  DB_DATABASE: "Dummy_data"
  ```
  ```sh
  DB_PORT: "8000"
  ```
  ```sh
  MYSQL_CONNECTION: "root:rootpassword@tcp(127.0.0.1:3306)/Dummy_data"
  ```

## Installation

The following terminal commands will allow you to run a local version of the server.

  ```sh
  docker build --tag docker-gs-ping .
  ```
  
  ```sh
  docker-compose up
  ```
  
  After the Docker container is running, Port 8080 should spin a MySQL server with Docker and Docker Compose (with Adminer). Port 8000 (http://localhost:8000/test) should render a neat User Interface. Port 8080 should now have the following login page:
  
  <img width="862" alt="Screen Shot 2021-08-04 at 2 57 24 PM" src="https://user-images.githubusercontent.com/70079800/128260862-7b7c2b6d-6a7a-42b0-a965-62c44d155aec.png">
  
  The username and password will be filled out, so clicking "login" will render a page with a list of databases. Instead of choosing an existing database, click "create database" on the top left page. Title the database "Dummy_data" The specific database in the main.go is referenced as "Dummy_data". The specific table in the "Dummy_data" database that is referenced throughout the code is titled "authors". Below, we have provided a screenshot of the table:
  

  <img width="862" alt="Screen Shot 2021-08-04 at 2 44 54 PM" src="https://user-images.githubusercontent.com/70079800/128259226-01c36260-5ed8-4c26-897a-05b2b24eaea6.png">
  
  The previous instructions should have provided you with information in regards to the creation of the database and table. The usage examples will further explain how the server can be utilized.
  
  
<!-- USAGE EXAMPLES -->
## Usage

In this section, please have the test.yml file open in an Integrated Development Environment. The YML file is where the database queries are located. It should be noted that each query can be found at http://localhost:8000/ after appending the query variable name to the slash. Examples are provided below.

A screenshot of this file is provided below:

<img width="561" alt="Screen Shot 2021-08-04 at 3 09 55 PM" src="https://user-images.githubusercontent.com/70079800/128261820-fb949050-cfd9-4550-bbf5-a1bebce6ac89.png">

The application can take multiple queries at once.

EXAMPLE:

Head to http://localhost:8000/test. This refers to the first "test" query. The database should render as a neat UI and look like the following:

<img width="1179" alt="Screen Shot 2021-08-04 at 3 13 49 PM" src="https://user-images.githubusercontent.com/70079800/128262197-b2ab6ab2-16f6-4ed4-bc8d-bdf81b1c95d9.png">

EXAMPLE 2:

Head to http://localhost:8000/test2. This refers to the "test2" query. The database should reder as a neat UI and look like the following:

<img width="1167" alt="Screen Shot 2021-08-04 at 3 14 32 PM" src="https://user-images.githubusercontent.com/70079800/128262261-ecac5f4d-d3b5-46a5-89fd-219927a50191.png">

The SQL queries in the YML file will be parsed in main.go and render on http://localhost:8000/

<!-- CONTACT -->
## Contact

Shivani Ahuja - shivaniahuja2001@gmail.com
<br />
Nick Almeder - [add email here]

Project Link: [https://github.com/ahujashivani/MySQL-Server](https://github.com/ahujashivani/MySQL-Server)

<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements
* [GitHub Emoji Cheat Sheet](https://www.webpagefx.com/tools/emoji-cheat-sheet)
* [Img Shields](https://shields.io)
* [Choose an Open Source License](https://choosealicense.com)
* [GitHub Pages](https://pages.github.com)
* [Animate.css](https://daneden.github.io/animate.css)
* [Loaders.css](https://connoratherton.com/loaders)
* [Slick Carousel](https://kenwheeler.github.io/slick)
* [Smooth Scroll](https://github.com/cferdinandi/smooth-scroll)
* [Sticky Kit](http://leafo.net/sticky-kit)
* [JVectorMap](http://jvectormap.com)
* [Font Awesome](https://fontawesome.com)





<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/othneildrew/Best-README-Template.svg?style=for-the-badge
[contributors-url]: https://github.com/othneildrew/Best-README-Template/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/othneildrew/Best-README-Template.svg?style=for-the-badge
[forks-url]: https://github.com/othneildrew/Best-README-Template/network/members
[stars-shield]: https://img.shields.io/github/stars/othneildrew/Best-README-Template.svg?style=for-the-badge
[stars-url]: https://github.com/othneildrew/Best-README-Template/stargazers
[issues-shield]: https://img.shields.io/github/issues/othneildrew/Best-README-Template.svg?style=for-the-badge
[issues-url]: https://github.com/othneildrew/Best-README-Template/issues
[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=for-the-badge
[license-url]: https://github.com/othneildrew/Best-README-Template/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/othneildrew
[product-screenshot]: images/screenshot.png
