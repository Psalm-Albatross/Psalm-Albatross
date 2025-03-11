# CIDR Calculation Tool

🌐 CIDR Calculation Tool helps with understanding and calculating CIDR ranges.

## Usage

1. **Calculate CIDR details:**
   ```sh
   $ ./psalm-cidr-tool calculate 192.168.1.0/24
   ```

2. **Split a CIDR block into smaller subnets:**
   ```sh
   $ ./psalm-cidr-tool split 192.168.1.0/24 26
   ```

3. **Merge multiple CIDR blocks:**
   ```sh
   $ ./psalm-cidr-tool merge 192.168.1.0/24 192.168.2.0/24
   ```

4. **Validate if an IP belongs to a CIDR block:**
   ```sh
   $ ./psalm-cidr-tool validate 192.168.1.1 192.168.1.0/24
   ```

## CIDR Basics

- 📘 CIDR notation defines IP ranges (e.g., 192.168.1.0/24).
- 📏 The prefix length (/24) determines how many addresses are available.
- 🛡️ The subnet mask controls how the IP space is divided.
- 🌐 Network Address: First IP in the range.
- 📡 Broadcast Address: Last IP in the range.
- 💻 Usable IPs: IPs available for devices (excluding network/broadcast IPs).

## Examples

- `/24` = 256 total IPs (254 usable)
- `/26` = 64 total IPs (62 usable)
- `/16` = 65,536 total IPs (65,534 usable)

## Commands

- `calculate <CIDR>`: Calculate details of a given CIDR block.
- `split <CIDR> <new prefix>`: Split a CIDR block into smaller subnets.
- `merge <CIDR1> <CIDR2> [CIDR3] ...`: Merge multiple CIDR blocks into a single range.
- `validate <IP> <CIDR>`: Validate if an IP belongs to a CIDR block.
- `conflict <CIDR1> <CIDR2>`: Check if two CIDRs conflict and show a valid CIDR range if they do.
- `generate <CIDR1> <CIDR2>`: Generate a new CIDR range that covers both input CIDRs.
- `help`: Print a detailed help page.
- `version`: Show the tool version.

## License

This project is licensed under the MIT License.
