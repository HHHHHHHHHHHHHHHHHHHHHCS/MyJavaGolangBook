package java_test;

public class HelloWorld{
    public static void main(String[] args){
        System.out.println("Hello, world!");
        for(String arg : args){
            System.out.println(arg);
        }
    }
}