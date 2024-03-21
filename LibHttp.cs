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
        public Http(String url, String method = "GET", string header = null)
        {
            WebRequest request = WebRequest.Create(url);
            request.Method = method;
            // add header
            if (method.ToUpper() == "POST") { 
                // set the content size
            }
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
            Http http = new Http("http://localhost:7777/app/web/system", "GET", "Content-Type: application/json\nage: 20");
            Console.WriteLine("{0}", http.getResponse());
            Console.ReadLine();
        }
    }
}
