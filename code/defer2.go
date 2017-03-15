f, err := os.Open("filename.ext")
if err != nil {
    log.Fatal(err)
}
defer f.Close()
// do something with the open *File f
