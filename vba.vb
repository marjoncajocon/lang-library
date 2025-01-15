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

Sub SendPostRequestXMLHTTP()
    Dim http As Object
    Dim url As String
    Dim data As String
    Dim responseText As String

    ' Set the URL and data for the request
    url = "https://your.api.endpoint/endpoint"
    data = "param1=value1&param2=value2" ' Data to be sent in POST request

    ' Create the XMLHTTP object
    Set http = CreateObject("MSXML2.XMLHTTP")

    ' Open the HTTP request
    http.Open "POST", url, False

    ' Set the required headers
    http.setRequestHeader "Content-Type", "application/x-www-form-urlencoded"

    ' Send the request with data
    http.send data

    ' Get the response text
    responseText = http.responseText

    ' Output the response
    Debug.Print responseText
    
    ' Clean up
    Set http = Nothing
End Sub
