 const files = sign_in.control.files;


  if (files.length <= 0) {
    alert('No file please select a file');
    return;
  }

  const file = files[0];
  if (file['type'] != 'image/png') {
    alert(`File must be: image/png, png file is required! not ${file.type}`);
    return;
  }

  // FileReader
  const reader = new FileReader();

  reader.onload = function (evt) {
      var binary = '';
      var bytes = new Uint8Array(reader.result);
      var len = bytes.byteLength;
      for (var i = 0; i < len; i++) {
          binary += String.fromCharCode(bytes[i]);
      }
      console.log(btoa(binary))
  }

  reader.readAsArrayBuffer(file)
