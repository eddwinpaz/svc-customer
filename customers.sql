-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: mysql-service:3306
-- Generation Time: May 09, 2020 at 12:52 AM
-- Server version: 5.7.30
-- PHP Version: 7.4.5

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `muble_db`
--

-- --------------------------------------------------------

--
-- Table structure for table `customers`
--

CREATE TABLE `customers` (
  `ID` int(11) NOT NULL,
  `customer_uuid` varchar(255) NOT NULL,
  `name` varchar(30) NOT NULL,
  `last_name` varchar(30) NOT NULL,
  `dni` varchar(20) DEFAULT 'DNI',
  `dni_type` enum('DNI','PASSPORT') NOT NULL,
  `email` varchar(100) NOT NULL,
  `phone` varchar(20) NOT NULL,
  `country` varchar(2) NOT NULL DEFAULT 'US',
  `password` varchar(150) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `customers`
--

INSERT INTO `customers` (`ID`, `customer_uuid`, `name`, `last_name`, `dni`, `dni_type`, `email`, `phone`, `country`, `password`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, '71f3a03f-ba84-4e31-a364-9a62f8e06993', 'User Name', 'User Lastname', '900000000', 'PASSPORT', 'fake1@gmail.com', '000000000', 'VE', 'fcc182b1296850192a241a4fcbed5f8cf33e3d6c', '2020-04-12 21:25:57', '2020-04-28 23:48:27', '2020-05-01 19:05:58'),
(2, '20946866-bf66-4873-aa5b-eb7ab7910721', 'User Name 2', 'User Lastname2', '700000000', 'DNI', 'fake2@gmail.com', '120000000', 'cl', 'b9b98302e9ec2146367058e67aee8ba964f36d60', '2020-04-12 21:26:25', '2020-05-01 18:23:29', '2020-05-01 19:06:50');
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
