-- MariaDB dump 10.19-11.2.2-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: tibiago
-- ------------------------------------------------------
-- Server version	11.2.2-MariaDB-1:11.2.2+maria~ubu2204

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Dumping data for table `bosses`
--

LOCK TABLES `bosses` WRITE;
/*!40000 ALTER TABLE `bosses` DISABLE KEYS */;
INSERT INTO `bosses` VALUES
('Ferumbras'),
('Oberon');
/*!40000 ALTER TABLE `bosses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `bosshunts`
--

LOCK TABLES `bosshunts` WRITE;
/*!40000 ALTER TABLE `bosshunts` DISABLE KEYS */;
INSERT INTO `bosshunts` VALUES
('2ca156a6-f546-4ee1-8e77-af433e45ed73','Ferumbras','Jaguna','2024-02-15 11:00:00.000'),
('f408e780-3b9e-45ea-8f92-b2622ebdc4c2','Ferumbras','Kendria','2024-02-15 19:00:00.000');
/*!40000 ALTER TABLE `bosshunts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `participants`
--

LOCK TABLES `participants` WRITE;
/*!40000 ALTER TABLE `participants` DISABLE KEYS */;
INSERT INTO `participants` VALUES
('099cc607-1050-4307-a5da-f1ffcebc12b5','f408e780-3b9e-45ea-8f92-b2622ebdc4c2','Filipem','knight','4029d401-921c-401f-bf47-0add5c1442f2'),
('afbef8e5-2e04-4c8e-9709-492880f92fa7','2ca156a6-f546-4ee1-8e77-af433e45ed73','Agresif Mob','sorcerer','4029d401-921c-401f-bf47-0add5c1442f2');
/*!40000 ALTER TABLE `participants` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES
('4029d401-921c-401f-bf47-0add5c1442f2');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `worlds`
--

LOCK TABLES `worlds` WRITE;
/*!40000 ALTER TABLE `worlds` DISABLE KEYS */;
INSERT INTO `worlds` VALUES
('Jaguna'),
('Kendria');
/*!40000 ALTER TABLE `worlds` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-02-15 17:21:54
