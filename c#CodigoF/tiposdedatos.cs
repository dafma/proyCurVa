using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Presentación1
{
    class Program
    {
        static void Main(string[] args)
        {
            int i = 0;
            decimal x = 0.0m;
            float f = 0.0f;
            double d = 0.0D;
            string cadena = "hola mundo";
            bool bandera = false;
            DateTime fecha = DateTime.MinValue;
            Console.WriteLine("EL valor de i es: {0}", i);
            Console.WriteLine("EL valor de x es: {0:C}", x);
            Console.WriteLine("EL valor de f es: {0:F2}", f);
            Console.WriteLine("EL valor de d es: {0:F2}", d);
            Console.WriteLine("EL valor de cadena es:"+ cadena);
            Console.WriteLine("EL valor de bandera es: "+ bandera.ToString());
            Console.WriteLine("EL valor de fecha es:" + fecha.ToShortDateString());
            Console.ReadKey();
        }
    }
}
