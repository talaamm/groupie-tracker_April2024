# Deployment Guide - Render.com

## Free Deployment on Render.com

This guide will help you deploy your Go application to Render.com for free.

### Prerequisites
- A GitHub account
- Your code pushed to a GitHub repository

### Step 1: Push to GitHub
1. Create a new repository on GitHub
2. Push your code to the repository:
```bash
git init
git add .
git commit -m "Initial commit"
git branch -M main
git remote add origin https://github.com/YOUR_USERNAME/YOUR_REPO_NAME.git
git push -u origin main
```

### Step 2: Deploy on Render
1. Go to [render.com](https://render.com) and sign up/login
2. Click "New +" and select "Web Service"
3. Connect your GitHub account and select your repository
4. Configure the service:
   - **Name**: groupie-tracker (or any name you prefer)
   - **Environment**: Docker
   - **Region**: Choose closest to your users
   - **Branch**: main
   - **Build Command**: `docker build -t groupie-tracker .`
   - **Start Command**: `docker run -p $PORT:405 groupie-tracker`
5. Click "Create Web Service"

### Step 3: Environment Variables
Render will automatically set the `PORT` environment variable, so no additional configuration is needed.

### Step 4: Wait for Deployment
- Render will build and deploy your application
- This usually takes 2-5 minutes
- You'll get a URL like: `https://your-app-name.onrender.com`

### Free Tier Limits
- 750 hours per month (enough for 24/7 uptime)
- Automatic sleep after 15 minutes of inactivity
- 512 MB RAM
- Shared CPU

### Troubleshooting
- Check the build logs if deployment fails
- Ensure all files are committed to GitHub
- Verify the Dockerfile is correct

Your app will be live at the provided URL once deployment is complete! 