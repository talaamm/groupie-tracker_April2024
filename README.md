# Groupie Tracker

The Groupie Tracker project involves building a user-friendly website that interfaces with a given API to display information about various music artists, their concert dates, and locations. This project focuses on the client-server interaction, handling API requests and rendering the data on the frontend.

## Overview

The project receives an API consisting of four parts:
1. **Artists**: Information about bands/artists including their name, image, formation year, and members.
2. **Locations**: Details of the artists' last and upcoming concert locations.
3. **Dates**: Information about the artists' last and upcoming concert dates.
4. **Relation**: Links between artists, concert dates, and locations.

The task is to create a website where users can visualize this information through various methods such as tables, lists, cards, or graphs. The project also includes creating events/actions that trigger server-side calls and displaying the results on the site.

## Features

1. **Data Display**:
   - The website presents artist information, concert dates, and locations in an engaging and clear format.
   - Uses data visualizations such as tables, cards, and graphics to enhance the user experience.

2. **Client-Server Communication**:
   - Implemented client-server interaction with HTTP requests/responses. 
   - The site triggers actions that communicate with the backend to fetch data dynamically.

3. **Error Handling**:
   - Ensures the site and server are stable, with proper error handling to prevent crashes.
   - All pages work correctly, and any errors are caught and handled gracefully.

4. **Backend (Go)**:
   - The backend is written in Go and handles the API requests, processes data, and serves it to the frontend.

5. **Frontend (HTML)**:
   - The frontend uses HTML to display the data fetched from the backend in a visually appealing format.

## Technologies Used
- **Backend**: Go
- **Frontend**: HTML
- **Client-Server**: RESTful API, HTTP requests and responses

## Instructions

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/talaamm/groupie-tracker_April2024.git
   ```

2. **Run the Backend**:
   - Navigate to the backend folder and run the Go server:
     ```bash
     go run main.go
     ```

3. **Run the Frontend**:
   - Open the HTML files in a browser to view the visualizations and interact with the website.

4. **API Usage**:
   - The site automatically fetches data from the API when a user interacts with it.

## Example Features

- **Artist Information**: Displays a list of bands/artists, their members, and images.
- **Concert Details**: Shows upcoming concert dates and locations for each artist.
- **Event Actions**: The site allows users to trigger specific actions (such as viewing upcoming concerts for a particular artist), which results in a dynamic data update.

## Error Handling

- The system is designed to ensure that errors are handled appropriately. The server and website will not crash unexpectedly, and the user will be provided with meaningful error messages in case of issues.

## Status
- The project has been successfully completed and passed all audits, including both the backend and frontend portions.
