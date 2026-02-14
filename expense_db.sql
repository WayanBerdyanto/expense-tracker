-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost:8889
-- Generation Time: Feb 14, 2026 at 11:10 AM
-- Server version: 8.0.35
-- PHP Version: 8.2.20

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `expense_db`
--

-- --------------------------------------------------------

--
-- Table structure for table `expenses`
--

CREATE TABLE `expenses` (
  `id` int NOT NULL,
  `description` varchar(255) NOT NULL,
  `amount` decimal(10,2) NOT NULL,
  `category` varchar(100) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `expenses`
--

INSERT INTO `expenses` (`id`, `description`, `amount`, `category`, `created_at`) VALUES
(2, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:10:40'),
(3, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:37:15'),
(4, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:37:27'),
(5, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:37:28'),
(6, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:37:28'),
(7, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:37:29'),
(8, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:37:29'),
(9, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:37:30'),
(10, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:37:30'),
(11, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:37:31'),
(12, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:37:31'),
(13, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:37:32'),
(14, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:37:32'),
(15, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:37:32'),
(16, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:37:34'),
(17, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:37:34'),
(18, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:37:35'),
(19, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:37:35'),
(20, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:37:36'),
(21, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:37:36'),
(22, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:37:37'),
(23, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:37:37'),
(24, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:37:38'),
(25, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:37:38'),
(26, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:37:38'),
(27, 'Lorem Ipsum', 2000.00, 'Food', '2026-02-07 14:37:39');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` enum('ADMIN','USER') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'USER',
  `updated_at` timestamp NOT NULL,
  `created_at` timestamp NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `password`, `role`, `updated_at`, `created_at`) VALUES
('5a20391d-6c66-4759-bdaf-798700d7acf9', 'Wayan Berdyanto', 'wayan.berdyanjto@gmail.com', '$2a$10$cyHZXn/dnzrA5eb43xgsz.bmwJ/jgePebSq0T0kHfj7iQV6U6G45e', 'USER', '2026-02-07 11:27:59', '2026-02-07 11:27:59'),
('9f052277-88d2-443f-aca1-0fe7cf0fae84', 'Wayan Berdyanto', 'wayan.berdyanto@gmail.com', '$2a$10$66AqZuwAyt23YtM2Dlrx..V2UZg9Hzgl5WuHT0twJbFeLzlmTDPYG', 'USER', '2026-02-07 11:24:20', '2026-02-07 11:24:20');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `expenses`
--
ALTER TABLE `expenses`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD KEY `email` (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `expenses`
--
ALTER TABLE `expenses`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=28;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
