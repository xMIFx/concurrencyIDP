package com.gitHub.xMIFx;

import org.junit.AfterClass;
import org.junit.BeforeClass;
import org.junit.Test;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class TestIslandTest {
	private static final Logger log = LoggerFactory.getLogger(TestIslandTest.class);
	private static int i = 0;
	@BeforeClass
	public static void oneTimeSetUp() {
		log.info("init");
		i++;
	}

	@Test
	public void testMethod() {
		log.info("test: " + i);
	}

	@AfterClass
	public static void oneTimeTearDown() {
		log.info("stop");
	}

}