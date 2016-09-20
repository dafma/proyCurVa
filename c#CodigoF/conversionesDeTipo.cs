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
            int i = 10;
            decimal x = 12.2m;
            bool bandera = false;
            string cadena = string.Empty;
            DateTime fecha = DateTime.MinValue;
            // conversion implicita por que no se necesita definir el var entero se puede converit a valor decimal
            x = i;
            //case conversiones demanera explicita
            i = (int)x;
            // otra forma e conversio
            i = Convert.ToInt32(x);
            Console.WriteLine("El valor de i es: {0}", i);
            Console.WriteLine("El valor de x es: {0:C}", x);
            //conversiones co metodos que establecen los metodos de conversion
            Console.WriteLine("El valor de bandera es" + bandera.ToString());
            Console.WriteLine("El valor de la cadena: " + cadena);
            Console.WriteLine("El valor de la fecha: " + fecha.ToShortDateString());
            Console.ReadKey();



        }
    }
}
