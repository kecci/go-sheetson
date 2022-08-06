# go-sheetson
This is Unofficially Library Golang API for https://sheetson.com/

## Introduction
Build Instant APIs from Google Sheets. 
Sheetson gives you an instant backend solution for all your coding projects. No setup, no worries, unlimited possibilities.

## Preparation
1. Register on https://sheetson.com/
2. Setup API Key on https://sheetson.com/console
3. Create spreadsheet on your google drive https://drive.google.com/
4. Share your spreadsheet with email `google@sheetson.com` as an editor.
5. Get SpreadsheetID from url, ex: https://docs.google.com/spreadsheets/d/1SOliBz08DFsNPxspR17VxI4Bb7UiDzPeftWfgmxB1SY/edit#gid=0 so your spreadsheetID is: `1SOliBz08DFsNPxspR17VxI4Bb7UiDzPeftWfgmxB1SY`
6. Get SheetName for example: `Sheet1`

### About sharing spreadsheet
Sheetson has taken security serious by not letting your spreadsheet published to the world. Instead, you only need to share your spreadsheet with our email `google@sheetson.com` as an editor, then we will handle the rest.

## API Features
- [x] Create - POST https://api.sheetson.com/v2/sheets/sheetName
- [ ] Read
  - [ ] List - GET https://api.sheetson.com/v2/sheets/sheetName
  - [ ] Row - GET https://api.sheetson.com/v2/sheets/sheetName/rowIndex
- [ ] Update - PATCH https://api.sheetson.com/v2/sheets/sheetName/rowIndex
- [ ] Delete - DELETE https://api.sheetson.com/v2/sheets/sheetName/rowIndex