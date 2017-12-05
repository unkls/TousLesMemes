-- MySQL dump 10.13  Distrib 5.7.17, for macos10.12 (x86_64)
--
-- Host: localhost    Database: memes
-- ------------------------------------------------------
-- Server version	5.6.35

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `CATEGORY`
--

DROP TABLE IF EXISTS `CATEGORY`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `CATEGORY` (
  `CAT.id` int(11) NOT NULL,
  `CAT.description` varchar(245) DEFAULT NULL,
  PRIMARY KEY (`CAT.id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `COMMENTS`
--

DROP TABLE IF EXISTS `COMMENTS`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `COMMENTS` (
  `COM.id` int(11) NOT NULL,
  `COM.user_id` int(11) NOT NULL,
  `COM.picture_id` int(11) NOT NULL,
  `COM.message` varchar(245) DEFAULT NULL,
  `COM.like` int(11) DEFAULT NULL,
  `COM.dislike` int(11) DEFAULT NULL,
  `COM.viewable` tinyint(4) NOT NULL DEFAULT '1',
  PRIMARY KEY (`COM.id`),
  KEY `FK_user_id_Comments_idx` (`COM.user_id`),
  KEY `FK_picture_id_Comments_idx` (`COM.picture_id`),
  CONSTRAINT `FK_pictures_id_Comments` FOREIGN KEY (`COM.picture_id`) REFERENCES `PICTURES` (`PCT.id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `FK_user_ID_Comments` FOREIGN KEY (`COM.user_id`) REFERENCES `USERS` (`user_id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `PASSWORD`
--

DROP TABLE IF EXISTS `PASSWORD`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `PASSWORD` (
  `PWD.id` int(11) NOT NULL,
  `PWD.password` varchar(45) DEFAULT NULL,
  `PWD.user_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`PWD.id`),
  KEY `foreignKey_user_id_idx` (`PWD.user_id`),
  CONSTRAINT `foreignKey_user_id` FOREIGN KEY (`PWD.user_id`) REFERENCES `USERS` (`user_id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `PICTURES`
--

DROP TABLE IF EXISTS `PICTURES`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `PICTURES` (
  `PCT.id` int(11) NOT NULL,
  `PCT.path` varchar(245) DEFAULT NULL,
  `PCT.id_category` int(11) DEFAULT NULL,
  `PCT.viewable` varchar(45) DEFAULT NULL,
  `PCT.user_id` int(11) DEFAULT NULL,
  `PCT.like` int(11) DEFAULT NULL,
  `PCT.dislike` int(11) DEFAULT NULL,
  PRIMARY KEY (`PCT.id`),
  KEY `ForeignKey_user_id_idx` (`PCT.user_id`),
  KEY `FK_id_categorie_idx` (`PCT.id_category`),
  CONSTRAINT `FK_id_categorie` FOREIGN KEY (`PCT.id_category`) REFERENCES `CATEGORY` (`CAT.id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `FK_user_id_pictures` FOREIGN KEY (`PCT.user_id`) REFERENCES `USERS` (`user_id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `USERS`
--

DROP TABLE IF EXISTS `USERS`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `USERS` (
  `user_id` int(11) NOT NULL,
  `user.nom` varchar(45) DEFAULT NULL,
  `user.prenom` varchar(45) DEFAULT NULL,
  `user.pseudo` varchar(45) DEFAULT NULL,
  `user.email` varchar(45) DEFAULT NULL,
  `user.age` int(3) DEFAULT NULL,
  `user.id_password` int(11) DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  KEY `ForeignKey_password_id_idx` (`user.id_password`),
  CONSTRAINT `ForeignKey_password_id` FOREIGN KEY (`user.id_password`) REFERENCES `PASSWORD` (`PWD.id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-11-10  9:54:45
