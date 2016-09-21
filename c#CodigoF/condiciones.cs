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
            //valores inicializados
            int valor1=0, valor2=0;
            string cadena;
            Console.WriteLine("Dame el primer valor");
            //almacenamos en una variable el valor obtenuido por el metodo readline
            cadena = Console.ReadLine();
            valor1 = Convert.ToInt32(cadena);
            Console.WriteLine("Dame el segundo valor");
            cadena = Console.ReadLine();
            //convertimos el valor de cadena en un valor entero
            valor2 = Convert.ToInt32(cadena);
            if(valor1 <= valor2)
                if (valor1 == valor2)
                    Console.WriteLine("EL valor 1 es igual al valor 2");
                else
                    Console.WriteLine("El valor 2 es mayor ");
            else
                Console.WriteLine("El valor 1 es mayor que el valor 2");
            Console.ReadKey();
        }
    }
}
