using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.IO;
using System.Net;

namespace Lib
{
    class Http
    {
        private string response = null;
        public Http(String url, String method = "GET", string header = null, string post_data = null)
        {
            WebRequest request = WebRequest.Create(url);
            request.Method = method;
            // add header
            if (header != null) {
                string [] h_arr = header.Split('\n');
                int len = h_arr.Length;
                for (int i = 0; i < len; i++) {
                   string [] h = h_arr[i].Split(':');
                   if (h.Length > 0) {
                       //Console.WriteLine("{0} - {1}", h[0].Trim(), h[1].Trim());
                       if (h[0] == "Content-Type")
                       {
                           request.ContentType = h[1].Trim();
                       }
                       else
                       {
                           request.Headers.Add(h[0].Trim(), h[1].Trim());
                       }
                   }
                }
            }
            // end add header

            // request postdata
            if (post_data != null)
            {
                // Convert the data to a byte array
                byte[] byteArray = Encoding.UTF8.GetBytes(post_data);

                // Set the content length header
                request.ContentLength = byteArray.Length;

                // Get the request stream and write the data to it
                using (Stream dataStream = request.GetRequestStream())
                {
                    dataStream.Write(byteArray, 0, byteArray.Length);
                }
            }
            // request end post data

            WebResponse resp = null;
            try
            {
                // Get the response
                resp = request.GetResponse();
                Stream dataStream = resp.GetResponseStream();
                StreamReader reader = new StreamReader(dataStream);
                this.response = reader.ReadToEnd();
            }
            catch (WebException e)
            {
                Console.WriteLine("WebException: {0}", e.Message);
            }
            finally {
                if (resp != null) {
                    resp.Close();
                }
            }
        }
        public string getResponse() {
            return this.response;
        }
    }

    class Program
    {
        static void Main(string[] args)
        {
            Http http = new Http("http://localhost:9000/info", "POST", "Content-Type: application/json", "hellow world this is sa sample");
            Console.WriteLine("{0}", http.getResponse());
            Console.ReadLine();
        }
    }
}
