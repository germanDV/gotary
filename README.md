# GOTARY

[logo](./frontend/src/assets/images/logo-universal.png)

## About

**Gotary** is a cross-platform desktop application that provides an easy way of cryptographically signing documents.

It produces a signature, represented as a hexadecimal string, that you save to a file and send together with the original file.
The receiving party can verify the signature using your public key and rest assured the file is exactly the one you signed.

Gotary lets you store contacts, as you need their public keys to verify signatures you receive from other people.


## Warning

This project is under heavy development, it should be considered _alpha_ software.

## Installation

**Work in progress**, binaries for Linux, MacOS and Windows will be available for download with each release.

## Develop

It's built with [wails](https://wails.io), using React for the UI.

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

To build a redistributable, production mode package, use `wails build`.
