-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Jun 06, 2021 at 01:43 PM
-- Server version: 10.4.19-MariaDB
-- PHP Version: 8.0.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `teric`
--

-- --------------------------------------------------------

--
-- Table structure for table `categories`
--

CREATE TABLE `categories` (
  `category_id` int(10) UNSIGNED NOT NULL,
  `category_name` varchar(255) DEFAULT NULL,
  `category_description` varchar(10000) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `categories`
--

INSERT INTO `categories` (`category_id`, `category_name`, `category_description`, `created_at`, `updated_at`) VALUES
(1, 'vegitables', 'Vegetable, in the broadest sense, any kind of plant life or plant product, namely “vegetable matter”; in common, narrow usage, the term vegetable usually refers to the fresh edible portions of certain herbaceous plants—roots, stems, leaves, flowers, fruit, or seeds', '2021-06-06 11:24:41', '2021-06-06 11:24:41'),
(2, 'fruits', 'fruit, the fleshy or dry ripened ovary of a flowering plant, enclosing the seed or seeds. Thus, apricots, bananas, and grapes, as well as bean pods, corn grains, tomatoes, cucumbers, and (in their shells) acorns and almonds, are all technically fruits', '2021-06-06 11:25:44', '2021-06-06 11:25:44');

-- --------------------------------------------------------

--
-- Table structure for table `category_products`
--

CREATE TABLE `category_products` (
  `id` int(10) UNSIGNED NOT NULL,
  `category_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `product_id` int(10) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `category_products`
--

INSERT INTO `category_products` (`id`, `category_id`, `created_at`, `updated_at`, `product_id`) VALUES
(1, 1, '2021-06-06 11:26:30', '2021-06-06 11:26:30', 1),
(2, 1, '2021-06-06 11:27:50', '2021-06-06 11:27:50', 2),
(3, 1, '2021-06-06 11:31:51', '2021-06-06 11:31:51', 3),
(4, 1, '2021-06-06 11:32:54', '2021-06-06 11:32:54', 4),
(5, 2, '2021-06-06 11:34:19', '2021-06-06 11:34:19', 5),
(6, 1, '2021-06-06 11:36:26', '2021-06-06 11:36:26', 6),
(7, 2, '2021-06-06 11:37:43', '2021-06-06 11:37:43', 7),
(8, 1, '2021-06-06 11:41:02', '2021-06-06 11:41:02', 8);

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `product_id` int(10) UNSIGNED NOT NULL,
  `product_name` varchar(255) DEFAULT NULL,
  `product_quantity` varchar(255) DEFAULT NULL,
  `product_description` varchar(10000) DEFAULT NULL,
  `product_price` int(10) UNSIGNED DEFAULT NULL,
  `sku` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`product_id`, `product_name`, `product_quantity`, `product_description`, `product_price`, `sku`, `created_at`, `updated_at`) VALUES
(1, 'Brinjal', '42', 'Brinjal is a rather small plant that grows up to 1.5 m. Brinjal is classified as a herb because of its non-woody stem. Its simple leaves are oblong to oval, slightly lobed, with its underside a paler green than the upper surface. Both leaves and stem are covered with fine hairs.', 30, 'KTPVTS', '2021-06-06 11:26:30', '2021-06-06 11:26:30'),
(2, 'Capsicum', '25', 'Capsicum, is also known as red pepper or chili pepper, is an herb. The fruit of the capsicum plant is used to make medicine. Capsicum is most commonly used for rheumatoid arthritis (RA), osteoarthritis, and other painful conditions.', 40, 'KTPVTS', '2021-06-06 11:27:50', '2021-06-06 11:27:50'),
(3, 'onion', '23', 'An onion is a round vegetable with a brown skin that grows underground. It has many white layers on its inside which have a strong, sharp smell and taste. ... It is made with fresh minced meat, cooked with onion and a rich tomato sauce.', 40, 'GJSKU', '2021-06-06 11:31:51', '2021-06-06 11:31:51'),
(4, 'tomato', '23', 'Tomato, (Solanum lycopersicum), flowering plant of the nightshade family (Solanaceae), cultivated extensively for its edible fruits. ... Labelled as a vegetable for nutritional purposes, tomatoes are a good source of vitamin C and the phytochemical lycopene', 18, 'JKIKU', '2021-06-06 11:32:54', '2021-06-06 11:32:54'),
(5, 'apple', '83', 'An apple is an edible fruit produced by an apple tree. Apple trees are cultivated worldwide and are the most widely grown species in the genus Malus. The tree originated in Central Asia, where its wild ancestor, Malus sieversii, is still found today.', 58, 'PIMKU', '2021-06-06 11:34:19', '2021-06-06 11:34:19'),
(6, 'Ginger', '53', 'Ginger, Zingiber officinale, is an erect, herbaceous perennial plant in the family Zingiberaceae grown for its edible rhizome (underground stem) which is widely used as a spice. The rhizome is brown, with a corky outer layer and pale-yellow scented center', 10, 'LAJKKU', '2021-06-06 11:36:26', '2021-06-06 11:36:26'),
(7, 'Banana', '30', 'A banana is a curved, yellow fruit with a thick skin and soft sweet flesh. If you eat a banana every day for breakfast, your roommate might nickname you the monkey. A banana is a tropical fruit that\'s quite popular all over the world. It grows in bunches on a banana tree', 40, 'AlJKKU', '2021-06-06 11:37:43', '2021-06-06 11:37:43'),
(8, 'Cabbage', '38', 'Smooth-leafed, firm-headed green cabbages are the most common, with smooth-leafed purple cabbages and crinkle-leafed savoy cabbages of both colours being rarer', 27, 'IUTUYT', '2021-06-06 11:41:02', '2021-06-06 11:41:02');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `username`, `password`, `created_at`) VALUES
(1, 'Bikash', 'bikashkumar01293@gmail.com', 'bikash12', '$2a$04$DaFf6NXRle92pXtkoL5dkueeqseRHQBE..mynyxRCYENdswgMuNtS', '2021-06-06 11:22:40');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`category_id`);

--
-- Indexes for table `category_products`
--
ALTER TABLE `category_products`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`product_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `categories`
--
ALTER TABLE `categories`
  MODIFY `category_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `category_products`
--
ALTER TABLE `category_products`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `product_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
