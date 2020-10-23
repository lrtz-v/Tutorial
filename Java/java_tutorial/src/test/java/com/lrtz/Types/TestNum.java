package com.lrtz.Types;

import org.junit.Assert;
import org.junit.Test;

/**
 * @author lvtao
 * @date 2020/10/23 21:47
 */
public class TestNum {

    @Test
    public void testNum() {
        int num1 = 1;
        Integer num2 = new Integer(1);
        Assert.assertTrue(num1 == num2);
    }

}
