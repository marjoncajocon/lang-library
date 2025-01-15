Sub SendHttpGETRequest()
    Dim oXMLHTTP As Object
    Set oXMLHTTP = CreateObject("MSXML2.XMLHTTP")
    
    ' Set the request URL
    Const URL As String = "https://www.google.com"
    
    ' Send the HTTP GET request
    oXMLHTTP.Open "GET", URL, False
    oXMLHTTP.Send
    
    ' Check if the request was successful
    If oXMLHTTP.Status = 200 Then
        MsgBox oXMLHTTP.responseText
    Else
        Debug.Print "Request Failed!"
        Debug.Print "Status Code: " & oXMLHTTP.Status
    End If
    
    ' Clean up
    Set oXMLHTTP = Nothing
End Sub
