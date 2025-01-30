function s2ab(s) {
      var buf = new ArrayBuffer(s.length);
      var view = new Uint8Array(buf);
      for (var i=0; i!=s.length; ++i) view[i] = s.charCodeAt(i) & 0xFF;
      return buf;
  }

  var blob = new Blob([s2ab(atob(data))], {
    type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet'
  });

  // const href = URL.createObjectURL(blob);
  // location.href = href;
  var link = document.createElement('a');
  link.href = window.URL.createObjectURL(blob);
  link.download = file_name + ".xlsx";
  link.click();
}
