import nodemailer from 'nodemailer';

const transporter = nodemailer.createTransport({
  service: 'gmail',
  auth: {
    user: process.env.MAIL_USER, // Your email address
    pass: process.env.MAIL_PASS, // Your app password
  },
});

export default transporter;
