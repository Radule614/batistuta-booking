﻿// <auto-generated />
using System;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Infrastructure;
using Microsoft.EntityFrameworkCore.Storage.ValueConversion;
using user_service.data.Db;

#nullable disable

namespace user_service.data.Migrations
{
    [DbContext(typeof(UserServiceDbContext))]
    partial class UserServiceDbContextModelSnapshot : ModelSnapshot
    {
        protected override void BuildModel(ModelBuilder modelBuilder)
        {
#pragma warning disable 612, 618
            modelBuilder
                .HasAnnotation("ProductVersion", "7.0.5")
                .HasAnnotation("Relational:MaxIdentifierLength", 64);

            modelBuilder.Entity("user_service.domain.Entities.User", b =>
                {
                    b.Property<Guid>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("char(36)");

                    b.Property<string>("Email")
                        .IsRequired()
                        .HasMaxLength(64)
                        .HasColumnType("varchar(64)");

                    b.Property<string>("FirstName")
                        .IsRequired()
                        .HasMaxLength(32)
                        .HasColumnType("varchar(32)");

                    b.Property<string>("LastName")
                        .IsRequired()
                        .HasMaxLength(32)
                        .HasColumnType("varchar(32)");

                    b.Property<string>("LivingPlace")
                        .IsRequired()
                        .HasMaxLength(64)
                        .HasColumnType("varchar(64)");

                    b.Property<string>("Password")
                        .IsRequired()
                        .HasMaxLength(256)
                        .HasColumnType("varchar(256)");

                    b.Property<string>("Role")
                        .IsRequired()
                        .HasMaxLength(32)
                        .HasColumnType("varchar(32)");

                    b.Property<string>("Username")
                        .IsRequired()
                        .HasMaxLength(64)
                        .HasColumnType("varchar(64)");

                    b.HasKey("Id");

                    b.HasAlternateKey("Email");

                    b.HasAlternateKey("Username");

                    b.ToTable("Users");

                    b.HasData(
                        new
                        {
                            Id = new Guid("422bd1e1-6d70-42cc-adc6-89a09b313c01"),
                            Email = "guest@test.com",
                            FirstName = "Guest",
                            LastName = "Guest",
                            LivingPlace = "Novi Sad, Srbija",
                            Password = "10000.GhMJYLVMJDUSKPYAt3G+oA==.2d2SyAT1CWcY/eNqJiKXKdTvjrWY2TftfJsHiOCy54g=",
                            Role = "Guest",
                            Username = "guest"
                        },
                        new
                        {
                            Id = new Guid("87abe34c-9935-44c4-aad5-af82f9442c77"),
                            Email = "host@test.com",
                            FirstName = "Host",
                            LastName = "Host",
                            LivingPlace = "Sabac, Srbija",
                            Password = "10000.pUrW8b1z1nt7+RFCVYWpWg==.jp+r7PJ49rgJwwZVAfLhCb2YMyCAZR3gXgrLnson2UQ=",
                            Role = "Host",
                            Username = "host"
                        });
                });
#pragma warning restore 612, 618
        }
    }
}
