package java_test;

public class TestString{
    public static void main(String[] args){
        String s1 = "abc1";
        String s2 = "abc1";
        System.out.println(s1 == s2);
        
        int x = 1;
        String s3 = "abc" + x;
        System.out.println(s1 == s3);
        
        s3 = s3.intern();
        System.out.println(111);
        
        TestString obj1 = new TestString();
        TestString obj2 = new TestString();
        System.out.println(obj1.hashCode());
        System.out.println(obj1.toString());
        System.out.println(obj1.equals(obj2));
        System.out.println(obj1.equals(obj1));

    }
}