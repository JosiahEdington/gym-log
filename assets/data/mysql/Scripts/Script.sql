CREATE TABLE `User` (
  `UserId` bigint,
  `FirstName` varchar(255),
  `LastName` varchar(255),
  `Email` varchar(255),
  `Username` varchar(255),
  `Weight` decimal(5,2),
  `DateOfBirth` date,
  `Sex` varchar(10),
  `CreatedDateTime` datetime,
  `CreatedBy` varchar(255),
  `UpdatedDateTime` datetime,
  `UpdatedBy` varchar(255),
  `IsDeleted` bit,
  `IsActive` bit,
  PRIMARY KEY (`UserId`)
);

CREATE TABLE `Workout` (
  `WorkoutId` bigint,
  `UserId` bigint,
  `StartedDateTime` datetime,
  `EndedDateTime` datetime,
  `Duration` decimal(4,2),
  `Category` varchar(255),
  `Type` varchar(255),
  `Name` varchar(255),
  `CreatedDateTime` datetime,
  `CreatedBy` varchar(255),
  `UpdatedDateTime` datetime,
  `UpdatedBy` varchar(255),
  `IsDeleted` bit,
  `Rank` int
);

CREATE TABLE `HealthMetric` (
  `HealthMetricId` bigint,
  `UserId` bigint,
  `WorkoutId` bigint,
  `MeasureTypeId` bigint,
  `ValueNumber` decimal(7,3),
  `ValueText` varchar(255),
  `MeasuredDateTime` datetime,
  `CreatedDateTime` datetime,
  `CreatedBy` varchar(255),
  `UpdatedDateTime` datetime,
  `UpdatedBy` varchar(255),
  `IsDeleted` bit
);

CREATE TABLE `SetId` (
  `SetId` bigint,
  `ExerciseId` bigint,
  `WorkoutId` bigint,
  `ExerciseType` bigint,
  `Order` int,
  `Reps` int,
  `Weight` int,
  `Duration` bigint,
  `Distance` bigint,
  `Pace` varchar(255),
  `RPE` int,
  `Rank` int,
  `StartedDateTime` datetime,
  `EndedDateTime` datetime,
  `CreatedDateTime` datetime,
  `CreatedBy` varchar(255),
  `UpdatedDateTime` datetime,
  `UpdatedBy` varchar(255),
  `IsDeleted` bit
);

CREATE TABLE `Exercise` (
  `ExerciseId` bigint,
  `WorkoutId` bigint,
  `Order` int,
  `ExerciseTypeId` bigint,
  `Weight` int,
  `Duration` bigint,
  `Distance` bigint,
  `Pace` varchar(255),
  `RPE` int,
  `Rank` int,
  `StartedDateTime` datetime,
  `EndedDateTime` datetime,
  `CreatedDateTime` datetime,
  `CreatedBy` varchar(255),
  `UpdatedDateTime` datetime,
  `UpdatedBy` varchar(255),
  `IsDeleted` bit
);

CREATE TABLE `MeasureType` (
  `MeasureTypeId` bigint,
  `MetricName` varchar(255),
  `Unit` varchar(255),
  `Category` varchar(255),
  `Type` varchar(255),
  `Description` varchar(510),
  `CreatedDateTime` datetime,
  `CreatedBy` varchar(255),
  `IsDeleted` bit
);

CREATE TABLE `ExerciseTypeId` (
  `ExerciseTypeId` bigint,
  `Name` varchar(255),
  `Unit` varchar(255),
  `MeasuredBy` varchar(255),
  `Description` varchar(510),
  `Type` varchar(255),
  `Category` varchar(255),
  `PrimaryTarget` varchar(255),
  `SecondaryTarget` varchar(255),
  `TertiaryTarget` varchar(255),
  `Equipment` varchar(255),
  `CreatedDateTime` datetime,
  `CreatedBy` varchar(255),
  `UpdatedDateTime` datetime,
  `UpdatedBy` varchar(255),
  `IsDeleted` bit
);

CREATE TABLE `Notes` (
  `NoteId` bigint,
  `UserId` bigint,
  `WorkoutId` bigint,
  `HealthMetricId` bigint,
  `ExerciseId` bigint,
  `SetRepId` bigint,
  `Rank` int,
  `Description` varchar(255),
  `Notes` varchar(510),
  `CreatedDateTime` datetime,
  `CreatedBy` varchar(255),
  `UpdatedDateTime` datetime,
  `UpdatedBy` varchar(255),
  `IsDeleted` bit
);

