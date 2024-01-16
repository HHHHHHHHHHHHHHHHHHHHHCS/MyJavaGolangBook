package java_test;

import java.util.ArrayList;
import java.util.List;

public class TestString implements Cloneable {
    public double pi = 3.14;

    @Override
    public TestString clone()
    {
        try {
            return (TestString) super.clone();
        } catch (CloneNotSupportedException e) {
            throw new RuntimeException(e);
        }
    }

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
        
        System.out.println("===========================");

        TestString c1 = new TestString();
        TestString c2 = c1.clone();
        TestString c3 = c1;
        c1.pi = 3.1415926;
        System.out.println(c1.pi);
        System.out.println(c2.pi);
        System.out.println(c3.pi);
        
        System.out.println("===========================");
        
        List<Integer> list = new ArrayList<>();
        list.add(1);
        list.add(2);
        list.add(3);
        System.out.println(list.toString());
        for(int item : list){
            System.out.println(item);
        }
    }
}