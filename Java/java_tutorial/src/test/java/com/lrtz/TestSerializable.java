package com.lrtz;

import org.junit.Test;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.ObjectOutputStream;
import java.io.Serializable;
import java.util.Arrays;

/**
 * @author lvtao
 * @date 2020/10/9 10:18
 */
public class TestSerializable {

    private class Point implements Serializable {
        Integer x;
        Integer y;
    }

    @Test
    public void testSerializable() {
        ByteArrayOutputStream buffer = new ByteArrayOutputStream();
        try (ObjectOutputStream output = new ObjectOutputStream(buffer)) {
            // 写入int:
            output.writeInt(12345);
            // 写入String:
            output.writeUTF("Hello");
            // 写入Object:
            output.writeObject(Double.valueOf(123.456));
        } catch (IOException e) {
            e.printStackTrace();
        }
        System.out.println(Arrays.toString(buffer.toByteArray()));
    }

}
