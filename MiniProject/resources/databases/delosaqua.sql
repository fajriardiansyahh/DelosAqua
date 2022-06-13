-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 13, 2022 at 06:41 PM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 7.4.29

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `delosaqua`
--

-- --------------------------------------------------------

--
-- Table structure for table `api`
--

CREATE TABLE `api` (
  `id` int(11) NOT NULL,
  `method` varchar(255) DEFAULT NULL,
  `host` varchar(255) DEFAULT NULL,
  `path` varchar(255) DEFAULT NULL,
  `ua` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `api`
--

INSERT INTO `api` (`id`, `method`, `host`, `path`, `ua`, `created_at`) VALUES
(1, 'GET', 'localhost:9090', '/api/v1/ponds', 'Thunder Client (https://www.thunderclient.com)', '2022-06-13 20:19:09'),
(2, 'POST', 'localhost:9090', '/api/v1/farms', 'Thunder Client (https://www.thunderclient.com)', '2022-06-13 20:19:30'),
(3, 'GET', 'localhost:9090', '/api/v1/farms', 'Thunder Client (https://www.thunderclient.com)', '2022-06-13 20:19:45'),
(4, 'PUT', 'localhost:9090', '/api/v1/farms', 'Thunder Client (https://www.thunderclient.com)', '2022-06-13 20:20:03'),
(5, 'POST', 'localhost:9090', '/api/v1/farms', 'Thunder Client (https://www.thunderclient.com)', '2022-06-13 20:20:18'),
(6, 'POST', 'localhost:9090', '/api/v1/farms', 'Thunder Client (https://www.thunderclient.com)', '2022-06-13 20:20:25'),
(7, 'PUT', 'localhost:9090', '/api/v1/farms', 'Thunder Client (https://www.thunderclient.com)', '2022-06-13 20:20:58'),
(8, 'DELETE', 'localhost:9090', '/api/v1/farms', 'Thunder Client (https://www.thunderclient.com)', '2022-06-13 20:21:08'),
(9, 'DELETE', 'localhost:9090', '/api/v1/farms', 'Thunder Client (https://www.thunderclient.com)', '2022-06-13 20:21:25'),
(10, 'PUT', 'localhost:9090', '/api/v1/farms', 'Thunder Client (https://www.thunderclient.com)', '2022-06-13 20:21:50'),
(11, 'GET', 'localhost:9090', '/api/v1/ponds', 'Thunder Client (https://www.thunderclient.com)', '2022-06-13 20:23:48'),
(12, 'POST', 'localhost:9090', '/api/v1/ponds', 'Thunder Client (https://www.thunderclient.com)', '2022-06-13 20:24:05'),
(13, 'POST', 'localhost:9090', '/api/v1/ponds', 'Thunder Client (https://www.thunderclient.com)', '2022-06-13 20:25:35'),
(14, 'PUT', 'localhost:9090', '/api/v1/ponds', 'Thunder Client (https://www.thunderclient.com)', '2022-06-13 20:26:08'),
(15, 'POST', 'localhost:9090', '/api/v1/ponds', 'Thunder Client (https://www.thunderclient.com)', '2022-06-13 20:26:52'),
(16, 'GET', 'localhost:9090', '/api/v1/ponds', 'Thunder Client (https://www.thunderclient.com)', '2022-06-13 20:26:56'),
(17, 'PUT', 'localhost:9090', '/api/v1/ponds', 'Thunder Client (https://www.thunderclient.com)', '2022-06-13 20:27:30'),
(18, 'DELETE', 'localhost:9090', '/api/v1/ponds', 'Thunder Client (https://www.thunderclient.com)', '2022-06-13 20:27:36');

-- --------------------------------------------------------

--
-- Table structure for table `farm`
--

CREATE TABLE `farm` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `thumbnails` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `farm`
--

INSERT INTO `farm` (`id`, `name`, `description`, `thumbnails`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'Updating Farm Name 1', 'Updating Description Farm Name 1', 'Updating Thumbnails Farm Name 1', '2022-06-13 20:19:30', '2022-06-13 13:21:50', NULL),
(2, 'Farm Name 2', 'Description Farm Name 2', 'Thumbnails Farm Name 2', '2022-06-13 20:20:18', NULL, NULL),
(3, 'Farm Name 3', 'Description Farm Name 3', 'Thumbnails Farm Name 3', '2022-06-13 20:20:25', NULL, '2022-06-13 13:21:25');

-- --------------------------------------------------------

--
-- Table structure for table `ponds`
--

CREATE TABLE `ponds` (
  `id` int(11) NOT NULL,
  `farm_id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `thumbnails` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `ponds`
--

INSERT INTO `ponds` (`id`, `farm_id`, `name`, `description`, `thumbnails`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 1, 'Ponds Name 1', 'Description Pond Name 1', 'Thumbnails Pond Name 1', '2022-06-13 20:25:35', '2022-06-13 13:26:08', NULL),
(2, 2, 'Updating Ponds Name 2', 'Updating Description Pond Name 2', 'Updating Thumbnails Pond Name 2', '2022-06-13 20:26:52', '2022-06-13 13:27:30', '2022-06-13 13:27:36');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `api`
--
ALTER TABLE `api`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `farm`
--
ALTER TABLE `farm`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `ponds`
--
ALTER TABLE `ponds`
  ADD PRIMARY KEY (`id`),
  ADD KEY `farm_id` (`farm_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `api`
--
ALTER TABLE `api`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;

--
-- AUTO_INCREMENT for table `farm`
--
ALTER TABLE `farm`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `ponds`
--
ALTER TABLE `ponds`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `ponds`
--
ALTER TABLE `ponds`
  ADD CONSTRAINT `ponds_ibfk_1` FOREIGN KEY (`farm_id`) REFERENCES `farm` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
